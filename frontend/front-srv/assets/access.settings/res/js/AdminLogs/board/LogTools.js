/*
 * Copyright 2007-2020 Charles du Jeu - Abstrium SAS <team (at) pyd.io>
 * This file is part of Pydio.
 *
 * Pydio is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Pydio is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Pydio.  If not, see <http://www.gnu.org/licenses/>.
 *
 * The latest code can be found at <https://pydio.com>.
 */

import Pydio from 'pydio'
import React from 'react'
import debounce from 'lodash.debounce'
import Log from '../model/Log'
import {RaisedButton, DatePicker, TimePicker, IconButton, FlatButton, FontIcon, IconMenu, MenuItem, Subheader, Dialog} from 'material-ui'
import {muiThemeable} from 'material-ui/styles'
const {ModernTextField, ModernSelectField, ModernStyles} = Pydio.requireLib('hoc');
const {moment} = Pydio.requireLib('boot');

class LogTools extends React.Component{

    constructor(props){
        super(props);
        this.state = {
            filter: "",
            filterMode: "fulltext",
            levelShow:false,
            serviceFilterShow: false,
        };
        this.publishStateChange = debounce(this.publishStateChange.bind(this), 250);
    }


    publishStateChange(){
        const {filter, serviceFilter, level, remoteAddress, userName, date, endDate} = this.state;
        const query = Log.buildQuery(filter, serviceFilter, level, remoteAddress, userName, date, endDate);
        this.props.onStateChange({query});
    }

    handleToggleShow(field){
        const fieldName = field + 'Show';
        const crt = this.state[fieldName];
        const s = {[fieldName]:!crt};
        if(crt){
           if(field === 'date' || field === 'endDate'){
               s['date'] = null;
               s['endDate'] = null;
               s['dateShow'] = false;
               s['endDateShow'] = false;
           } else {
               s[field] = null;
           }
           s['page'] = 0;
        } else if(field === 'date' && this.state.endDateShow){
            s['endDate'] = null;
            s['endDateShow'] = false;
        }
        this.setState(s, this.publishStateChange.bind(this));
    }

    handleFilterChange(val, keyName) {
        this.setState({[keyName]: val, page: 0}, this.publishStateChange.bind(this))
    }

    handleDateChange(date, time = null) {
        if(time){
            date.setHours(time.getHours(), time.getMinutes());
        }
        const {endDate, endDateShow} = this.state;
        if(endDateShow && !endDate && date !== undefined){
            let end = new Date();
            end.setHours(23, 59, 59);
            this.setState({endDate: end})
        }
        this.setState({date: date, page: 0}, this.publishStateChange.bind(this));
    }

    handleEndDateChange(date, time = null) {
        if(time){
            date.setHours(time.getHours(), time.getMinutes());
        }
        this.setState({endDate: date, page: 0}, this.publishStateChange.bind(this));
    }

    handleExport(format) {
        const {filter, serviceFilter, level, remoteAddress, userName, date, endDate} = this.state;
        const {service} = this.props;
        const dateString = (date? date.getFullYear() + '-' + (date.getMonth()+1) + '-' + date.getDate() : '');
        const query = Log.buildQuery(filter, serviceFilter, level, remoteAddress, userName, date, endDate);
        Log.downloadLogs(service || 'sys', query, format).then(blob => {
            const url = window.URL.createObjectURL(blob);
            let filename = 'cells-logs-';
            if (dateString){
                filename += dateString;
            } else {
                filename += 'filtered';
            }
            filename += '.' + format.toLowerCase();
            if(navigator.userAgent.toLowerCase().indexOf('firefox') > -1) {
                this.setState({
                    exportUrl:url,
                    exportFilename:filename,
                    exportOnClick:()=>{
                        setTimeout(()=>{
                            window.URL.revokeObjectURL(url);
                        }, 100);
                        this.setState({exportUrl:null, exportFilename: null});
                    }
                });
                return;
            }
            const link = document.createElement('a');
            link.href = url;
            link.download = filename;
            link.click();
            setTimeout(() => {
                window.URL.revokeObjectURL(url);
            }, 100);
        });
    }

    render(){
        const {pydio, disableExport, muiTheme, focus} = this.props;
        const adminStyles = AdminComponents.AdminStyles(muiTheme.palette);
        const focusBadge = {
            backgroundColor: '#FBE9E7',
            height: 35,
            lineHeight: '35px',
            fontSize: 15,
            padding: '0 10px',
            marginRight: 5,
            color: '#FF5722',
            borderRadius: 3
        };

        const {filter, date, dateShow, endDate, endDateShow, serviceFilter, serviceFilterShow, level, levelShow, userName, userNameShow, remoteAddress, remoteAddressShow, exportUrl, exportFilename, exportOnClick} = this.state;
        const {MessageHash} = pydio;
        const hasFilter = filter || serviceFilter || date || endDate || level || userName || remoteAddress;
        const checkIcon = <FontIcon style={{top: 0}} className={"mdi mdi-check"}/>;
        return (
            <div style={{display: 'flex', alignItems: 'center', width: '100%', marginTop: 3}}>

                {focus &&
                    <div style={focusBadge}>Focus on +/- 5 minutes at {moment(new Date(focus*1000)).format('hh:mm:ss')}</div>
                }

                <div style={{marginRight: 5, width: 170}} >
                    <ModernTextField hintText={MessageHash["ajxp_admin.logs.3"]} onChange={(e, v) => this.handleFilterChange(v, 'filter')} fullWidth={true}/>
                </div>

                {levelShow &&
                    <div style={{marginRight: 5, marginTop: -2, width: 100}}>
                        <ModernSelectField hintText={"Level"} fullWidth={true} value={level}
                                           onChange={(e, i, v) => this.handleFilterChange(v, 'level')}>
                            <MenuItem primaryText={""}/>
                            <MenuItem primaryText={"ERROR"} value={"ERROR"}/>
                            <MenuItem primaryText={"INFO"} value={"INFO"}/>
                            <MenuItem primaryText={"DEBUG"} value={"DEBUG"}/>
                        </ModernSelectField>
                    </div>
                }

                {serviceFilterShow &&
                    <div style={{marginRight: 5, width: 80}} >
                        <ModernTextField hintText={"Service"} fullWidth={true} value={serviceFilter} onChange={(e, v) => this.handleFilterChange(v, 'serviceFilter')} />
                    </div>
                }

                {remoteAddressShow &&
                    <div style={{marginRight: 5, width: 80}} >
                        <ModernTextField hintText={"IP"} fullWidth={true} onChange={(e, v) => this.handleFilterChange(v, 'remoteAddress')} />
                    </div>
                }

                {userNameShow &&
                    <div style={{marginRight: 5, width: 80}} >
                        <ModernTextField hintText={"Login"} fullWidth={true} onChange={(e, v) => this.handleFilterChange(v, 'userName')} />
                    </div>
                }

                {dateShow && !endDateShow &&
                <div style={{display:'flex', alignItems:'center'}}>
                    <DatePicker hintText={MessageHash["ajxp_admin.logs.2"]} onChange={(e, date) => this.handleDateChange(date)}
                                autoOk={true} maxDate={new Date()} value={date}
                                showYearSelector={true} style={{width: 120}} textFieldStyle={{width: 120}} {...ModernStyles.textField}/>
                    <IconButton iconClassName={"mdi mdi-close"} tooltip={"Clear"} onTouchTap={() => {this.handleDateChange(undefined)}} {...adminStyles.props.header.iconButton}/>
                </div>
                }
                {endDateShow &&
                <div style={{display:'flex', alignItems:'center'}}>
                    <DatePicker hintText={'From'} onChange={(e, date) => this.handleDateChange(date)}
                                autoOk={true} maxDate={new Date()} value={date}
                                showYearSelector={true} style={{width: 100}} textFieldStyle={{width: 96}} {...ModernStyles.textField} />
                    <TimePicker hintText={'at...'} disabled={!date} onChange={(e, time) => this.handleDateChange(date, time)}
                                autoOk={true} value={date}
                                style={{width: 100}} textFieldStyle={{width: 96}} {...ModernStyles.textField} />
                    <DatePicker hintText={'To'} onChange={(e, date) => this.handleEndDateChange(date)}
                                autoOk={true} minDate={this.state.date} maxDate={new Date()} value={endDate}
                                showYearSelector={true} style={{width: 100}} textFieldStyle={{width: 96}} {...ModernStyles.textField} />
                    <TimePicker hintText={'at...'} disabled={!endDate} onChange={(e, time) => this.handleEndDateChange(endDate, time)}
                                autoOk={true} value={endDate}
                                style={{width: 100}} textFieldStyle={{width: 96}} {...ModernStyles.textField} />
                    <IconButton iconClassName={"mdi mdi-close"} tooltip={"Clear"} onTouchTap={() => {this.handleDateChange(undefined); this.handleEndDateChange(undefined)}} {...adminStyles.props.header.iconButton}/>
                </div>
                }
                <IconMenu
                    iconButtonElement={<IconButton iconClassName={"mdi mdi-filter-variant"} tooltip={MessageHash['ajxp_admin.logs.3']} {...adminStyles.props.header.iconButton}/>}
                    anchorOrigin={{vertical:'top', horizontal:'right'}}
                    targetOrigin={{vertical:'top', horizontal:'right'}}
                    desktop={true}
                >
                    {<Subheader>{MessageHash['ajxp_admin.logs.filter.legend']}</Subheader>}
                    <MenuItem primaryText={MessageHash['ajxp_admin.logs.2']}  rightIcon={dateShow && !endDateShow ? checkIcon : null} onTouchTap={()=>{this.handleToggleShow('date')}}/>
                    <MenuItem primaryText={MessageHash['ajxp_admin.logs.filter.period']}  rightIcon={endDateShow ? checkIcon : null} onTouchTap={()=>{this.handleToggleShow('endDate')}}/>
                    <MenuItem primaryText={"Level"}  rightIcon={levelShow ? checkIcon : null} onTouchTap={()=>{this.handleToggleShow('level')}}/>
                    <MenuItem primaryText={"Service"}  rightIcon={serviceFilterShow ? checkIcon : null} onTouchTap={()=>{this.handleToggleShow('serviceFilter')}}/>
                    <MenuItem primaryText={"User Login"}  rightIcon={userNameShow ? checkIcon : null} onTouchTap={()=>{this.handleToggleShow('userName')}}/>
                    <MenuItem primaryText={"IP"}  rightIcon={remoteAddressShow ? checkIcon : null} onTouchTap={()=>{this.handleToggleShow('remoteAddress')}}/>
                </IconMenu>

                {!disableExport &&
                    <IconMenu
                        iconButtonElement={<IconButton iconClassName={"mdi mdi-download"} tooltip={MessageHash["ajxp_admin.logs.11"]} {...adminStyles.props.header.iconButton}/>}
                        anchorOrigin={{vertical:'top', horizontal:'right'}}
                        targetOrigin={{vertical:'top', horizontal:'right'}}
                        desktop={true}
                    >
                        {!hasFilter && <Subheader>{MessageHash['ajxp_admin.logs.export.disabled']}</Subheader>}
                        {hasFilter && <Subheader>{MessageHash['ajxp_admin.logs.11']}</Subheader>}
                        <MenuItem primaryText="CSV" rightIcon={<FontIcon style={{top: 0}} className={"mdi mdi-file-delimited"}/>} onTouchTap={()=>{this.handleExport('CSV')}} disabled={!hasFilter} />
                        <MenuItem primaryText="XLSX" rightIcon={<FontIcon style={{top: 0}} className={"mdi mdi-file-excel"}/>} onTouchTap={()=>{this.handleExport('XLSX')}} disabled={!hasFilter} />
                        {exportUrl && <Subheader><a href={exportUrl} download={exportFilename}>{exportFilename}</a></Subheader>}
                    </IconMenu>
                }
                <Dialog
                    open={!!exportUrl}
                    modal={true}
                    title={MessageHash['ajxp_admin.logs.11']}
                    actions={[<FlatButton label={"Cancel"} onTouchTap={exportOnClick}/>]}
                >
                    <span style={{fontSize:13}}>
                        {MessageHash['ajxp_admin.logs.export.clicklink']}: <a style={{textDecoration:'underline'}} href={exportUrl} download={exportFilename} onClick={exportOnClick}>{exportFilename}</a>
                    </span>
                </Dialog>

            </div>
        )
    }


}
LogTools = muiThemeable()(LogTools);
export {LogTools as default}