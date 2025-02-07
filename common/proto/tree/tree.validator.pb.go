// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tree.proto

package tree

import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ReadNodeRequest) Validate() error {
	if this.Node != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Node); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Node", err)
		}
	}
	return nil
}
func (this *ReadNodeResponse) Validate() error {
	if this.Node != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Node); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Node", err)
		}
	}
	return nil
}
func (this *ListNodesRequest) Validate() error {
	if this.Node != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Node); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Node", err)
		}
	}
	return nil
}
func (this *ListNodesResponse) Validate() error {
	if this.Node != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Node); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Node", err)
		}
	}
	return nil
}
func (this *WrappingStreamerResponse) Validate() error {
	if oneOfNester, ok := this.GetData().(*WrappingStreamerResponse_ListNodesResponse); ok {
		if oneOfNester.ListNodesResponse != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.ListNodesResponse); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ListNodesResponse", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*WrappingStreamerResponse_NodeChangeEvent); ok {
		if oneOfNester.NodeChangeEvent != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.NodeChangeEvent); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("NodeChangeEvent", err)
			}
		}
	}
	return nil
}
func (this *CreateNodeRequest) Validate() error {
	if this.Node != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Node); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Node", err)
		}
	}
	return nil
}
func (this *CreateNodeResponse) Validate() error {
	if this.Node != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Node); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Node", err)
		}
	}
	return nil
}
func (this *UpdateNodeRequest) Validate() error {
	if this.From != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.From); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("From", err)
		}
	}
	if this.To != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.To); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("To", err)
		}
	}
	return nil
}
func (this *UpdateNodeResponse) Validate() error {
	if this.Node != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Node); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Node", err)
		}
	}
	return nil
}
func (this *DeleteNodeRequest) Validate() error {
	if this.Node != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Node); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Node", err)
		}
	}
	return nil
}
func (this *DeleteNodeResponse) Validate() error {
	return nil
}
func (this *IndexationSession) Validate() error {
	if this.RootNode != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RootNode); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RootNode", err)
		}
	}
	return nil
}
func (this *IndexationOperation) Validate() error {
	return nil
}
func (this *OpenSessionRequest) Validate() error {
	if this.Session != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Session); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Session", err)
		}
	}
	return nil
}
func (this *OpenSessionResponse) Validate() error {
	if this.Session != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Session); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Session", err)
		}
	}
	return nil
}
func (this *FlushSessionRequest) Validate() error {
	if this.Session != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Session); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Session", err)
		}
	}
	return nil
}
func (this *FlushSessionResponse) Validate() error {
	if this.Session != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Session); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Session", err)
		}
	}
	return nil
}
func (this *CloseSessionRequest) Validate() error {
	if this.Session != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Session); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Session", err)
		}
	}
	return nil
}
func (this *CloseSessionResponse) Validate() error {
	if this.Session != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Session); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Session", err)
		}
	}
	return nil
}
func (this *WatchNodeRequest) Validate() error {
	if this.Node != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Node); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Node", err)
		}
	}
	return nil
}
func (this *WatchNodeResponse) Validate() error {
	if this.Node != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Node); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Node", err)
		}
	}
	return nil
}
func (this *SearchRequest) Validate() error {
	if this.Query != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Query); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Query", err)
		}
	}
	return nil
}
func (this *SearchFacet) Validate() error {
	return nil
}
func (this *SearchResponse) Validate() error {
	if this.Node != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Node); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Node", err)
		}
	}
	if this.Facet != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Facet); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Facet", err)
		}
	}
	return nil
}
func (this *CreateVersionRequest) Validate() error {
	if this.Node != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Node); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Node", err)
		}
	}
	if this.TriggerEvent != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TriggerEvent); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TriggerEvent", err)
		}
	}
	return nil
}
func (this *CreateVersionResponse) Validate() error {
	if this.Version != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Version); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Version", err)
		}
	}
	return nil
}
func (this *ListVersionsRequest) Validate() error {
	if this.Node != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Node); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Node", err)
		}
	}
	return nil
}
func (this *ListVersionsResponse) Validate() error {
	if this.Version != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Version); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Version", err)
		}
	}
	return nil
}
func (this *HeadVersionRequest) Validate() error {
	if this.Node != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Node); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Node", err)
		}
	}
	return nil
}
func (this *HeadVersionResponse) Validate() error {
	if this.Version != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Version); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Version", err)
		}
	}
	return nil
}
func (this *StoreVersionRequest) Validate() error {
	if this.Node != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Node); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Node", err)
		}
	}
	if this.Version != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Version); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Version", err)
		}
	}
	return nil
}
func (this *StoreVersionResponse) Validate() error {
	for _, item := range this.PruneVersions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PruneVersions", err)
			}
		}
	}
	return nil
}
func (this *PruneVersionsRequest) Validate() error {
	if this.UniqueNode != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UniqueNode); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UniqueNode", err)
		}
	}
	return nil
}
func (this *PruneVersionsResponse) Validate() error {
	return nil
}
func (this *VersioningPolicy) Validate() error {
	for _, item := range this.KeepPeriods {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("KeepPeriods", err)
			}
		}
	}
	return nil
}
func (this *VersioningKeepPeriod) Validate() error {
	return nil
}
func (this *Node) Validate() error {
	for _, item := range this.Commits {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Commits", err)
			}
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	for _, item := range this.AppearsIn {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AppearsIn", err)
			}
		}
	}
	return nil
}
func (this *WorkspaceRelativePath) Validate() error {
	return nil
}
func (this *ChangeLog) Validate() error {
	if this.Event != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Event); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Event", err)
		}
	}
	return nil
}
func (this *Query) Validate() error {
	if this.GeoQuery != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.GeoQuery); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("GeoQuery", err)
		}
	}
	return nil
}
func (this *GeoQuery) Validate() error {
	if this.Center != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Center); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Center", err)
		}
	}
	if this.TopLeft != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TopLeft); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TopLeft", err)
		}
	}
	if this.BottomRight != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.BottomRight); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("BottomRight", err)
		}
	}
	return nil
}
func (this *GeoPoint) Validate() error {
	return nil
}
func (this *StreamChangesRequest) Validate() error {
	return nil
}
func (this *NodeChangeEvent) Validate() error {
	if this.Source != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Source); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Source", err)
		}
	}
	if this.Target != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Target); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Target", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *IndexEvent) Validate() error {
	return nil
}
func (this *GetEncryptionKeyRequest) Validate() error {
	if this.Node != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Node); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Node", err)
		}
	}
	return nil
}
func (this *GetEncryptionKeyResponse) Validate() error {
	return nil
}
func (this *SyncChange) Validate() error {
	if this.Node != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Node); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Node", err)
		}
	}
	return nil
}
func (this *SyncChangeNode) Validate() error {
	return nil
}
func (this *PutSyncChangeResponse) Validate() error {
	return nil
}
func (this *SearchSyncChangeRequest) Validate() error {
	return nil
}
