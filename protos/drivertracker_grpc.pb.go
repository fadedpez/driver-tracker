// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protos

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DriverTrackerAPIClient is the client API for DriverTrackerAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DriverTrackerAPIClient interface {
	StoreDriver(ctx context.Context, in *StoreDriverRequest, opts ...grpc.CallOption) (*StoreDriverResponse, error)
	GetDriver(ctx context.Context, in *GetDriverRequest, opts ...grpc.CallOption) (*GetDriverResponse, error)
	SearchDriver(ctx context.Context, in *SearchDriverRequest, opts ...grpc.CallOption) (*SearchDriverResponse, error)
	StoreTeam(ctx context.Context, in *StoreTeamRequest, opts ...grpc.CallOption) (*StoreTeamResponse, error)
	SearchTeamByName(ctx context.Context, in *SearchTeamByNameRequest, opts ...grpc.CallOption) (*SearchTeamByNameResponse, error)
	GetTeam(ctx context.Context, in *GetTeamRequest, opts ...grpc.CallOption) (*GetTeamResponse, error)
	StoreGrandPrix(ctx context.Context, in *StoreGrandPrixRequest, opts ...grpc.CallOption) (*StoreGrandPrixResponse, error)
	GetGrandPrix(ctx context.Context, in *GetGrandPrixRequest, opts ...grpc.CallOption) (*GetGrandPrixResponse, error)
	SearchGrandPrix(ctx context.Context, in *SearchGrandPrixRequest, opts ...grpc.CallOption) (*SearchGrandPrixResponse, error)
	StoreQualifying(ctx context.Context, in *StoreQualifyingRequest, opts ...grpc.CallOption) (*StoreQualifyingResponse, error)
	GetQualifying(ctx context.Context, in *GetQualifyingRequest, opts ...grpc.CallOption) (*GetQualifyingResponse, error)
	SearchQualifying(ctx context.Context, in *SearchQualifyingRequest, opts ...grpc.CallOption) (*SearchQualifyingResponse, error)
	StoreTrack(ctx context.Context, in *StoreTrackRequest, opts ...grpc.CallOption) (*StoreTrackResponse, error)
	GetTrack(ctx context.Context, in *GetTrackRequest, opts ...grpc.CallOption) (*GetTrackResponse, error)
	SearchTrack(ctx context.Context, in *SearchTrackRequest, opts ...grpc.CallOption) (*SearchTrackResponse, error)
	StoreRound(ctx context.Context, in *StoreRoundRequest, opts ...grpc.CallOption) (*StoreRoundResponse, error)
	GetRound(ctx context.Context, in *GetRoundRequest, opts ...grpc.CallOption) (*GetRoundResponse, error)
	SearchRound(ctx context.Context, in *SearchRoundRequest, opts ...grpc.CallOption) (*SearchRoundResponse, error)
	StoreSeason(ctx context.Context, in *StoreSeasonRequest, opts ...grpc.CallOption) (*StoreSeasonResponse, error)
	GetSeason(ctx context.Context, in *GetSeasonRequest, opts ...grpc.CallOption) (*GetSeasonResponse, error)
	SearchSeason(ctx context.Context, in *SearchSeasonRequest, opts ...grpc.CallOption) (*SearchSeasonResponse, error)
}

type driverTrackerAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewDriverTrackerAPIClient(cc grpc.ClientConnInterface) DriverTrackerAPIClient {
	return &driverTrackerAPIClient{cc}
}

func (c *driverTrackerAPIClient) StoreDriver(ctx context.Context, in *StoreDriverRequest, opts ...grpc.CallOption) (*StoreDriverResponse, error) {
	out := new(StoreDriverResponse)
	err := c.cc.Invoke(ctx, "/drivertracker.api.alpha.DriverTrackerAPI/StoreDriver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverTrackerAPIClient) GetDriver(ctx context.Context, in *GetDriverRequest, opts ...grpc.CallOption) (*GetDriverResponse, error) {
	out := new(GetDriverResponse)
	err := c.cc.Invoke(ctx, "/drivertracker.api.alpha.DriverTrackerAPI/GetDriver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverTrackerAPIClient) SearchDriver(ctx context.Context, in *SearchDriverRequest, opts ...grpc.CallOption) (*SearchDriverResponse, error) {
	out := new(SearchDriverResponse)
	err := c.cc.Invoke(ctx, "/drivertracker.api.alpha.DriverTrackerAPI/SearchDriver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverTrackerAPIClient) StoreTeam(ctx context.Context, in *StoreTeamRequest, opts ...grpc.CallOption) (*StoreTeamResponse, error) {
	out := new(StoreTeamResponse)
	err := c.cc.Invoke(ctx, "/drivertracker.api.alpha.DriverTrackerAPI/StoreTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverTrackerAPIClient) SearchTeamByName(ctx context.Context, in *SearchTeamByNameRequest, opts ...grpc.CallOption) (*SearchTeamByNameResponse, error) {
	out := new(SearchTeamByNameResponse)
	err := c.cc.Invoke(ctx, "/drivertracker.api.alpha.DriverTrackerAPI/SearchTeamByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverTrackerAPIClient) GetTeam(ctx context.Context, in *GetTeamRequest, opts ...grpc.CallOption) (*GetTeamResponse, error) {
	out := new(GetTeamResponse)
	err := c.cc.Invoke(ctx, "/drivertracker.api.alpha.DriverTrackerAPI/GetTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverTrackerAPIClient) StoreGrandPrix(ctx context.Context, in *StoreGrandPrixRequest, opts ...grpc.CallOption) (*StoreGrandPrixResponse, error) {
	out := new(StoreGrandPrixResponse)
	err := c.cc.Invoke(ctx, "/drivertracker.api.alpha.DriverTrackerAPI/StoreGrandPrix", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverTrackerAPIClient) GetGrandPrix(ctx context.Context, in *GetGrandPrixRequest, opts ...grpc.CallOption) (*GetGrandPrixResponse, error) {
	out := new(GetGrandPrixResponse)
	err := c.cc.Invoke(ctx, "/drivertracker.api.alpha.DriverTrackerAPI/GetGrandPrix", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverTrackerAPIClient) SearchGrandPrix(ctx context.Context, in *SearchGrandPrixRequest, opts ...grpc.CallOption) (*SearchGrandPrixResponse, error) {
	out := new(SearchGrandPrixResponse)
	err := c.cc.Invoke(ctx, "/drivertracker.api.alpha.DriverTrackerAPI/SearchGrandPrix", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverTrackerAPIClient) StoreQualifying(ctx context.Context, in *StoreQualifyingRequest, opts ...grpc.CallOption) (*StoreQualifyingResponse, error) {
	out := new(StoreQualifyingResponse)
	err := c.cc.Invoke(ctx, "/drivertracker.api.alpha.DriverTrackerAPI/StoreQualifying", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverTrackerAPIClient) GetQualifying(ctx context.Context, in *GetQualifyingRequest, opts ...grpc.CallOption) (*GetQualifyingResponse, error) {
	out := new(GetQualifyingResponse)
	err := c.cc.Invoke(ctx, "/drivertracker.api.alpha.DriverTrackerAPI/GetQualifying", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverTrackerAPIClient) SearchQualifying(ctx context.Context, in *SearchQualifyingRequest, opts ...grpc.CallOption) (*SearchQualifyingResponse, error) {
	out := new(SearchQualifyingResponse)
	err := c.cc.Invoke(ctx, "/drivertracker.api.alpha.DriverTrackerAPI/SearchQualifying", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverTrackerAPIClient) StoreTrack(ctx context.Context, in *StoreTrackRequest, opts ...grpc.CallOption) (*StoreTrackResponse, error) {
	out := new(StoreTrackResponse)
	err := c.cc.Invoke(ctx, "/drivertracker.api.alpha.DriverTrackerAPI/StoreTrack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverTrackerAPIClient) GetTrack(ctx context.Context, in *GetTrackRequest, opts ...grpc.CallOption) (*GetTrackResponse, error) {
	out := new(GetTrackResponse)
	err := c.cc.Invoke(ctx, "/drivertracker.api.alpha.DriverTrackerAPI/GetTrack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverTrackerAPIClient) SearchTrack(ctx context.Context, in *SearchTrackRequest, opts ...grpc.CallOption) (*SearchTrackResponse, error) {
	out := new(SearchTrackResponse)
	err := c.cc.Invoke(ctx, "/drivertracker.api.alpha.DriverTrackerAPI/SearchTrack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverTrackerAPIClient) StoreRound(ctx context.Context, in *StoreRoundRequest, opts ...grpc.CallOption) (*StoreRoundResponse, error) {
	out := new(StoreRoundResponse)
	err := c.cc.Invoke(ctx, "/drivertracker.api.alpha.DriverTrackerAPI/StoreRound", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverTrackerAPIClient) GetRound(ctx context.Context, in *GetRoundRequest, opts ...grpc.CallOption) (*GetRoundResponse, error) {
	out := new(GetRoundResponse)
	err := c.cc.Invoke(ctx, "/drivertracker.api.alpha.DriverTrackerAPI/GetRound", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverTrackerAPIClient) SearchRound(ctx context.Context, in *SearchRoundRequest, opts ...grpc.CallOption) (*SearchRoundResponse, error) {
	out := new(SearchRoundResponse)
	err := c.cc.Invoke(ctx, "/drivertracker.api.alpha.DriverTrackerAPI/SearchRound", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverTrackerAPIClient) StoreSeason(ctx context.Context, in *StoreSeasonRequest, opts ...grpc.CallOption) (*StoreSeasonResponse, error) {
	out := new(StoreSeasonResponse)
	err := c.cc.Invoke(ctx, "/drivertracker.api.alpha.DriverTrackerAPI/StoreSeason", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverTrackerAPIClient) GetSeason(ctx context.Context, in *GetSeasonRequest, opts ...grpc.CallOption) (*GetSeasonResponse, error) {
	out := new(GetSeasonResponse)
	err := c.cc.Invoke(ctx, "/drivertracker.api.alpha.DriverTrackerAPI/GetSeason", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverTrackerAPIClient) SearchSeason(ctx context.Context, in *SearchSeasonRequest, opts ...grpc.CallOption) (*SearchSeasonResponse, error) {
	out := new(SearchSeasonResponse)
	err := c.cc.Invoke(ctx, "/drivertracker.api.alpha.DriverTrackerAPI/SearchSeason", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriverTrackerAPIServer is the server API for DriverTrackerAPI service.
// All implementations must embed UnimplementedDriverTrackerAPIServer
// for forward compatibility
type DriverTrackerAPIServer interface {
	StoreDriver(context.Context, *StoreDriverRequest) (*StoreDriverResponse, error)
	GetDriver(context.Context, *GetDriverRequest) (*GetDriverResponse, error)
	SearchDriver(context.Context, *SearchDriverRequest) (*SearchDriverResponse, error)
	StoreTeam(context.Context, *StoreTeamRequest) (*StoreTeamResponse, error)
	SearchTeamByName(context.Context, *SearchTeamByNameRequest) (*SearchTeamByNameResponse, error)
	GetTeam(context.Context, *GetTeamRequest) (*GetTeamResponse, error)
	StoreGrandPrix(context.Context, *StoreGrandPrixRequest) (*StoreGrandPrixResponse, error)
	GetGrandPrix(context.Context, *GetGrandPrixRequest) (*GetGrandPrixResponse, error)
	SearchGrandPrix(context.Context, *SearchGrandPrixRequest) (*SearchGrandPrixResponse, error)
	StoreQualifying(context.Context, *StoreQualifyingRequest) (*StoreQualifyingResponse, error)
	GetQualifying(context.Context, *GetQualifyingRequest) (*GetQualifyingResponse, error)
	SearchQualifying(context.Context, *SearchQualifyingRequest) (*SearchQualifyingResponse, error)
	StoreTrack(context.Context, *StoreTrackRequest) (*StoreTrackResponse, error)
	GetTrack(context.Context, *GetTrackRequest) (*GetTrackResponse, error)
	SearchTrack(context.Context, *SearchTrackRequest) (*SearchTrackResponse, error)
	StoreRound(context.Context, *StoreRoundRequest) (*StoreRoundResponse, error)
	GetRound(context.Context, *GetRoundRequest) (*GetRoundResponse, error)
	SearchRound(context.Context, *SearchRoundRequest) (*SearchRoundResponse, error)
	StoreSeason(context.Context, *StoreSeasonRequest) (*StoreSeasonResponse, error)
	GetSeason(context.Context, *GetSeasonRequest) (*GetSeasonResponse, error)
	SearchSeason(context.Context, *SearchSeasonRequest) (*SearchSeasonResponse, error)
}

// UnimplementedDriverTrackerAPIServer must be embedded to have forward compatible implementations.
type UnimplementedDriverTrackerAPIServer struct {
}

func (UnimplementedDriverTrackerAPIServer) StoreDriver(context.Context, *StoreDriverRequest) (*StoreDriverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreDriver not implemented")
}
func (UnimplementedDriverTrackerAPIServer) GetDriver(context.Context, *GetDriverRequest) (*GetDriverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDriver not implemented")
}
func (UnimplementedDriverTrackerAPIServer) SearchDriver(context.Context, *SearchDriverRequest) (*SearchDriverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchDriver not implemented")
}
func (UnimplementedDriverTrackerAPIServer) StoreTeam(context.Context, *StoreTeamRequest) (*StoreTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreTeam not implemented")
}
func (UnimplementedDriverTrackerAPIServer) SearchTeamByName(context.Context, *SearchTeamByNameRequest) (*SearchTeamByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchTeamByName not implemented")
}
func (UnimplementedDriverTrackerAPIServer) GetTeam(context.Context, *GetTeamRequest) (*GetTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeam not implemented")
}
func (UnimplementedDriverTrackerAPIServer) StoreGrandPrix(context.Context, *StoreGrandPrixRequest) (*StoreGrandPrixResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreGrandPrix not implemented")
}
func (UnimplementedDriverTrackerAPIServer) GetGrandPrix(context.Context, *GetGrandPrixRequest) (*GetGrandPrixResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGrandPrix not implemented")
}
func (UnimplementedDriverTrackerAPIServer) SearchGrandPrix(context.Context, *SearchGrandPrixRequest) (*SearchGrandPrixResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchGrandPrix not implemented")
}
func (UnimplementedDriverTrackerAPIServer) StoreQualifying(context.Context, *StoreQualifyingRequest) (*StoreQualifyingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreQualifying not implemented")
}
func (UnimplementedDriverTrackerAPIServer) GetQualifying(context.Context, *GetQualifyingRequest) (*GetQualifyingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQualifying not implemented")
}
func (UnimplementedDriverTrackerAPIServer) SearchQualifying(context.Context, *SearchQualifyingRequest) (*SearchQualifyingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchQualifying not implemented")
}
func (UnimplementedDriverTrackerAPIServer) StoreTrack(context.Context, *StoreTrackRequest) (*StoreTrackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreTrack not implemented")
}
func (UnimplementedDriverTrackerAPIServer) GetTrack(context.Context, *GetTrackRequest) (*GetTrackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrack not implemented")
}
func (UnimplementedDriverTrackerAPIServer) SearchTrack(context.Context, *SearchTrackRequest) (*SearchTrackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchTrack not implemented")
}
func (UnimplementedDriverTrackerAPIServer) StoreRound(context.Context, *StoreRoundRequest) (*StoreRoundResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreRound not implemented")
}
func (UnimplementedDriverTrackerAPIServer) GetRound(context.Context, *GetRoundRequest) (*GetRoundResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRound not implemented")
}
func (UnimplementedDriverTrackerAPIServer) SearchRound(context.Context, *SearchRoundRequest) (*SearchRoundResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchRound not implemented")
}
func (UnimplementedDriverTrackerAPIServer) StoreSeason(context.Context, *StoreSeasonRequest) (*StoreSeasonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreSeason not implemented")
}
func (UnimplementedDriverTrackerAPIServer) GetSeason(context.Context, *GetSeasonRequest) (*GetSeasonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSeason not implemented")
}
func (UnimplementedDriverTrackerAPIServer) SearchSeason(context.Context, *SearchSeasonRequest) (*SearchSeasonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchSeason not implemented")
}
func (UnimplementedDriverTrackerAPIServer) mustEmbedUnimplementedDriverTrackerAPIServer() {}

// UnsafeDriverTrackerAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DriverTrackerAPIServer will
// result in compilation errors.
type UnsafeDriverTrackerAPIServer interface {
	mustEmbedUnimplementedDriverTrackerAPIServer()
}

func RegisterDriverTrackerAPIServer(s grpc.ServiceRegistrar, srv DriverTrackerAPIServer) {
	s.RegisterService(&DriverTrackerAPI_ServiceDesc, srv)
}

func _DriverTrackerAPI_StoreDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreDriverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverTrackerAPIServer).StoreDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drivertracker.api.alpha.DriverTrackerAPI/StoreDriver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverTrackerAPIServer).StoreDriver(ctx, req.(*StoreDriverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverTrackerAPI_GetDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDriverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverTrackerAPIServer).GetDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drivertracker.api.alpha.DriverTrackerAPI/GetDriver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverTrackerAPIServer).GetDriver(ctx, req.(*GetDriverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverTrackerAPI_SearchDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchDriverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverTrackerAPIServer).SearchDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drivertracker.api.alpha.DriverTrackerAPI/SearchDriver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverTrackerAPIServer).SearchDriver(ctx, req.(*SearchDriverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverTrackerAPI_StoreTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverTrackerAPIServer).StoreTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drivertracker.api.alpha.DriverTrackerAPI/StoreTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverTrackerAPIServer).StoreTeam(ctx, req.(*StoreTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverTrackerAPI_SearchTeamByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchTeamByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverTrackerAPIServer).SearchTeamByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drivertracker.api.alpha.DriverTrackerAPI/SearchTeamByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverTrackerAPIServer).SearchTeamByName(ctx, req.(*SearchTeamByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverTrackerAPI_GetTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverTrackerAPIServer).GetTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drivertracker.api.alpha.DriverTrackerAPI/GetTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverTrackerAPIServer).GetTeam(ctx, req.(*GetTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverTrackerAPI_StoreGrandPrix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreGrandPrixRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverTrackerAPIServer).StoreGrandPrix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drivertracker.api.alpha.DriverTrackerAPI/StoreGrandPrix",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverTrackerAPIServer).StoreGrandPrix(ctx, req.(*StoreGrandPrixRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverTrackerAPI_GetGrandPrix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGrandPrixRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverTrackerAPIServer).GetGrandPrix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drivertracker.api.alpha.DriverTrackerAPI/GetGrandPrix",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverTrackerAPIServer).GetGrandPrix(ctx, req.(*GetGrandPrixRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverTrackerAPI_SearchGrandPrix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchGrandPrixRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverTrackerAPIServer).SearchGrandPrix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drivertracker.api.alpha.DriverTrackerAPI/SearchGrandPrix",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverTrackerAPIServer).SearchGrandPrix(ctx, req.(*SearchGrandPrixRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverTrackerAPI_StoreQualifying_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreQualifyingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverTrackerAPIServer).StoreQualifying(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drivertracker.api.alpha.DriverTrackerAPI/StoreQualifying",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverTrackerAPIServer).StoreQualifying(ctx, req.(*StoreQualifyingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverTrackerAPI_GetQualifying_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQualifyingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverTrackerAPIServer).GetQualifying(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drivertracker.api.alpha.DriverTrackerAPI/GetQualifying",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverTrackerAPIServer).GetQualifying(ctx, req.(*GetQualifyingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverTrackerAPI_SearchQualifying_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchQualifyingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverTrackerAPIServer).SearchQualifying(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drivertracker.api.alpha.DriverTrackerAPI/SearchQualifying",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverTrackerAPIServer).SearchQualifying(ctx, req.(*SearchQualifyingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverTrackerAPI_StoreTrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreTrackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverTrackerAPIServer).StoreTrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drivertracker.api.alpha.DriverTrackerAPI/StoreTrack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverTrackerAPIServer).StoreTrack(ctx, req.(*StoreTrackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverTrackerAPI_GetTrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTrackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverTrackerAPIServer).GetTrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drivertracker.api.alpha.DriverTrackerAPI/GetTrack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverTrackerAPIServer).GetTrack(ctx, req.(*GetTrackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverTrackerAPI_SearchTrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchTrackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverTrackerAPIServer).SearchTrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drivertracker.api.alpha.DriverTrackerAPI/SearchTrack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverTrackerAPIServer).SearchTrack(ctx, req.(*SearchTrackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverTrackerAPI_StoreRound_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRoundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverTrackerAPIServer).StoreRound(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drivertracker.api.alpha.DriverTrackerAPI/StoreRound",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverTrackerAPIServer).StoreRound(ctx, req.(*StoreRoundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverTrackerAPI_GetRound_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverTrackerAPIServer).GetRound(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drivertracker.api.alpha.DriverTrackerAPI/GetRound",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverTrackerAPIServer).GetRound(ctx, req.(*GetRoundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverTrackerAPI_SearchRound_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRoundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverTrackerAPIServer).SearchRound(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drivertracker.api.alpha.DriverTrackerAPI/SearchRound",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverTrackerAPIServer).SearchRound(ctx, req.(*SearchRoundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverTrackerAPI_StoreSeason_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreSeasonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverTrackerAPIServer).StoreSeason(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drivertracker.api.alpha.DriverTrackerAPI/StoreSeason",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverTrackerAPIServer).StoreSeason(ctx, req.(*StoreSeasonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverTrackerAPI_GetSeason_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSeasonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverTrackerAPIServer).GetSeason(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drivertracker.api.alpha.DriverTrackerAPI/GetSeason",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverTrackerAPIServer).GetSeason(ctx, req.(*GetSeasonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverTrackerAPI_SearchSeason_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchSeasonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverTrackerAPIServer).SearchSeason(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drivertracker.api.alpha.DriverTrackerAPI/SearchSeason",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverTrackerAPIServer).SearchSeason(ctx, req.(*SearchSeasonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DriverTrackerAPI_ServiceDesc is the grpc.ServiceDesc for DriverTrackerAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DriverTrackerAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "drivertracker.api.alpha.DriverTrackerAPI",
	HandlerType: (*DriverTrackerAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreDriver",
			Handler:    _DriverTrackerAPI_StoreDriver_Handler,
		},
		{
			MethodName: "GetDriver",
			Handler:    _DriverTrackerAPI_GetDriver_Handler,
		},
		{
			MethodName: "SearchDriver",
			Handler:    _DriverTrackerAPI_SearchDriver_Handler,
		},
		{
			MethodName: "StoreTeam",
			Handler:    _DriverTrackerAPI_StoreTeam_Handler,
		},
		{
			MethodName: "SearchTeamByName",
			Handler:    _DriverTrackerAPI_SearchTeamByName_Handler,
		},
		{
			MethodName: "GetTeam",
			Handler:    _DriverTrackerAPI_GetTeam_Handler,
		},
		{
			MethodName: "StoreGrandPrix",
			Handler:    _DriverTrackerAPI_StoreGrandPrix_Handler,
		},
		{
			MethodName: "GetGrandPrix",
			Handler:    _DriverTrackerAPI_GetGrandPrix_Handler,
		},
		{
			MethodName: "SearchGrandPrix",
			Handler:    _DriverTrackerAPI_SearchGrandPrix_Handler,
		},
		{
			MethodName: "StoreQualifying",
			Handler:    _DriverTrackerAPI_StoreQualifying_Handler,
		},
		{
			MethodName: "GetQualifying",
			Handler:    _DriverTrackerAPI_GetQualifying_Handler,
		},
		{
			MethodName: "SearchQualifying",
			Handler:    _DriverTrackerAPI_SearchQualifying_Handler,
		},
		{
			MethodName: "StoreTrack",
			Handler:    _DriverTrackerAPI_StoreTrack_Handler,
		},
		{
			MethodName: "GetTrack",
			Handler:    _DriverTrackerAPI_GetTrack_Handler,
		},
		{
			MethodName: "SearchTrack",
			Handler:    _DriverTrackerAPI_SearchTrack_Handler,
		},
		{
			MethodName: "StoreRound",
			Handler:    _DriverTrackerAPI_StoreRound_Handler,
		},
		{
			MethodName: "GetRound",
			Handler:    _DriverTrackerAPI_GetRound_Handler,
		},
		{
			MethodName: "SearchRound",
			Handler:    _DriverTrackerAPI_SearchRound_Handler,
		},
		{
			MethodName: "StoreSeason",
			Handler:    _DriverTrackerAPI_StoreSeason_Handler,
		},
		{
			MethodName: "GetSeason",
			Handler:    _DriverTrackerAPI_GetSeason_Handler,
		},
		{
			MethodName: "SearchSeason",
			Handler:    _DriverTrackerAPI_SearchSeason_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "drivertracker.proto",
}
