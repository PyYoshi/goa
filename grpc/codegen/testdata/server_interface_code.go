package testdata

const UnaryRPCsServerInterfaceCode = `// MethodUnaryRPCA implements the "MethodUnaryRPCA" method in
// pb.ServiceUnaryRPCsServer interface.
func (s *Server) MethodUnaryRPCA(ctx context.Context, message *pb.MethodUnaryRPCARequest) (*pb.MethodUnaryRPCAResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "MethodUnaryRPCA")
	ctx = context.WithValue(ctx, goa.ServiceKey, "ServiceUnaryRPCs")
	resp, err := s.MethodUnaryRPCAH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*pb.MethodUnaryRPCAResponse), nil
}

// MethodUnaryRPCB implements the "MethodUnaryRPCB" method in
// pb.ServiceUnaryRPCsServer interface.
func (s *Server) MethodUnaryRPCB(ctx context.Context, message *pb.MethodUnaryRPCBRequest) (*pb.MethodUnaryRPCBResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "MethodUnaryRPCB")
	ctx = context.WithValue(ctx, goa.ServiceKey, "ServiceUnaryRPCs")
	resp, err := s.MethodUnaryRPCBH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*pb.MethodUnaryRPCBResponse), nil
}
`

const UnaryRPCNoPayloadServerInterfaceCode = `// MethodUnaryRPCNoPayload implements the "MethodUnaryRPCNoPayload" method in
// pb.ServiceUnaryRPCNoPayloadServer interface.
func (s *Server) MethodUnaryRPCNoPayload(ctx context.Context, message *pb.MethodUnaryRPCNoPayloadRequest) (*pb.MethodUnaryRPCNoPayloadResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "MethodUnaryRPCNoPayload")
	ctx = context.WithValue(ctx, goa.ServiceKey, "ServiceUnaryRPCNoPayload")
	resp, err := s.MethodUnaryRPCNoPayloadH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*pb.MethodUnaryRPCNoPayloadResponse), nil
}
`

const UnaryRPCNoResultServerInterfaceCode = `// MethodUnaryRPCNoResult implements the "MethodUnaryRPCNoResult" method in
// pb.ServiceUnaryRPCNoResultServer interface.
func (s *Server) MethodUnaryRPCNoResult(ctx context.Context, message *pb.MethodUnaryRPCNoResultRequest) (*pb.MethodUnaryRPCNoResultResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "MethodUnaryRPCNoResult")
	ctx = context.WithValue(ctx, goa.ServiceKey, "ServiceUnaryRPCNoResult")
	resp, err := s.MethodUnaryRPCNoResultH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*pb.MethodUnaryRPCNoResultResponse), nil
}
`

const UnaryRPCWithErrorsServerInterfaceCode = `// MethodUnaryRPCWithErrors implements the "MethodUnaryRPCWithErrors" method in
// pb.ServiceUnaryRPCWithErrorsServer interface.
func (s *Server) MethodUnaryRPCWithErrors(ctx context.Context, message *pb.MethodUnaryRPCWithErrorsRequest) (*pb.MethodUnaryRPCWithErrorsResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "MethodUnaryRPCWithErrors")
	ctx = context.WithValue(ctx, goa.ServiceKey, "ServiceUnaryRPCWithErrors")
	resp, err := s.MethodUnaryRPCWithErrorsH.Handle(ctx, message)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "timeout":
				return nil, goagrpc.NewStatusError(codes.Canceled, err)
			case "internal":
				return nil, goagrpc.NewStatusError(codes.Unknown, err)
			case "bad_request":
				return nil, goagrpc.NewStatusError(codes.InvalidArgument, err)
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*pb.MethodUnaryRPCWithErrorsResponse), nil
}
`

const ServerStreamingRPCServerInterfaceCode = `// MethodServerStreamingRPC implements the "MethodServerStreamingRPC" method in
// pb.ServiceServerStreamingRPCServer interface.
func (s *Server) MethodServerStreamingRPC(message *pb.MethodServerStreamingRPCRequest, stream pb.ServiceServerStreamingRPC_MethodServerStreamingRPCServer) error {
	ctx := stream.Context()
	ctx = context.WithValue(ctx, goa.MethodKey, "MethodServerStreamingRPC")
	ctx = context.WithValue(ctx, goa.ServiceKey, "ServiceServerStreamingRPC")
	p, err := s.MethodServerStreamingRPCH.Decode(ctx, message)
	if err != nil {
		return goagrpc.EncodeError(err)
	}
	ep := &serviceserverstreamingrpc.MethodServerStreamingRPCEndpointInput{
		Stream:  &MethodServerStreamingRPCServerStream{stream: stream},
		Payload: p.(int),
	}
	err = s.MethodServerStreamingRPCH.Handle(ctx, ep)
	if err != nil {
		return goagrpc.EncodeError(err)
	}
	return nil
}
`

const ClientStreamingRPCServerInterfaceCode = `// MethodClientStreamingRPC implements the "MethodClientStreamingRPC" method in
// pb.ServiceClientStreamingRPCServer interface.
func (s *Server) MethodClientStreamingRPC(stream pb.ServiceClientStreamingRPC_MethodClientStreamingRPCServer) error {
	ctx := stream.Context()
	ctx = context.WithValue(ctx, goa.MethodKey, "MethodClientStreamingRPC")
	ctx = context.WithValue(ctx, goa.ServiceKey, "ServiceClientStreamingRPC")
	p, err := s.MethodClientStreamingRPCH.Decode(ctx, nil)
	if err != nil {
		return goagrpc.EncodeError(err)
	}
	ep := &serviceclientstreamingrpc.MethodClientStreamingRPCEndpointInput{
		Stream: &MethodClientStreamingRPCServerStream{stream: stream},
	}
	err = s.MethodClientStreamingRPCH.Handle(ctx, ep)
	if err != nil {
		return goagrpc.EncodeError(err)
	}
	return nil
}
`

const ClientStreamingRPCWithPayloadServerInterfaceCode = `// MethodClientStreamingRPCWithPayload implements the
// "MethodClientStreamingRPCWithPayload" method in
// pb.ServiceClientStreamingRPCWithPayloadServer interface.
func (s *Server) MethodClientStreamingRPCWithPayload(stream pb.ServiceClientStreamingRPCWithPayload_MethodClientStreamingRPCWithPayloadServer) error {
	ctx := stream.Context()
	ctx = context.WithValue(ctx, goa.MethodKey, "MethodClientStreamingRPCWithPayload")
	ctx = context.WithValue(ctx, goa.ServiceKey, "ServiceClientStreamingRPCWithPayload")
	p, err := s.MethodClientStreamingRPCWithPayloadH.Decode(ctx, nil)
	if err != nil {
		return goagrpc.EncodeError(err)
	}
	ep := &serviceclientstreamingrpcwithpayload.MethodClientStreamingRPCWithPayloadEndpointInput{
		Stream:  &MethodClientStreamingRPCWithPayloadServerStream{stream: stream},
		Payload: p.(int),
	}
	err = s.MethodClientStreamingRPCWithPayloadH.Handle(ctx, ep)
	if err != nil {
		return goagrpc.EncodeError(err)
	}
	return nil
}
`

const BidirectionalStreamingRPCServerInterfaceCode = `// MethodBidirectionalStreamingRPC implements the
// "MethodBidirectionalStreamingRPC" method in
// pb.ServiceBidirectionalStreamingRPCServer interface.
func (s *Server) MethodBidirectionalStreamingRPC(stream pb.ServiceBidirectionalStreamingRPC_MethodBidirectionalStreamingRPCServer) error {
	ctx := stream.Context()
	ctx = context.WithValue(ctx, goa.MethodKey, "MethodBidirectionalStreamingRPC")
	ctx = context.WithValue(ctx, goa.ServiceKey, "ServiceBidirectionalStreamingRPC")
	p, err := s.MethodBidirectionalStreamingRPCH.Decode(ctx, nil)
	if err != nil {
		return goagrpc.EncodeError(err)
	}
	ep := &servicebidirectionalstreamingrpc.MethodBidirectionalStreamingRPCEndpointInput{
		Stream: &MethodBidirectionalStreamingRPCServerStream{stream: stream},
	}
	err = s.MethodBidirectionalStreamingRPCH.Handle(ctx, ep)
	if err != nil {
		return goagrpc.EncodeError(err)
	}
	return nil
}
`

const BidirectionalStreamingRPCWithPayloadServerInterfaceCode = `// MethodBidirectionalStreamingRPCWithPayload implements the
// "MethodBidirectionalStreamingRPCWithPayload" method in
// pb.ServiceBidirectionalStreamingRPCWithPayloadServer interface.
func (s *Server) MethodBidirectionalStreamingRPCWithPayload(stream pb.ServiceBidirectionalStreamingRPCWithPayload_MethodBidirectionalStreamingRPCWithPayloadServer) error {
	ctx := stream.Context()
	ctx = context.WithValue(ctx, goa.MethodKey, "MethodBidirectionalStreamingRPCWithPayload")
	ctx = context.WithValue(ctx, goa.ServiceKey, "ServiceBidirectionalStreamingRPCWithPayload")
	p, err := s.MethodBidirectionalStreamingRPCWithPayloadH.Decode(ctx, nil)
	if err != nil {
		return goagrpc.EncodeError(err)
	}
	ep := &servicebidirectionalstreamingrpcwithpayload.MethodBidirectionalStreamingRPCWithPayloadEndpointInput{
		Stream:  &MethodBidirectionalStreamingRPCWithPayloadServerStream{stream: stream},
		Payload: p.(*servicebidirectionalstreamingrpcwithpayload.Payload),
	}
	err = s.MethodBidirectionalStreamingRPCWithPayloadH.Handle(ctx, ep)
	if err != nil {
		return goagrpc.EncodeError(err)
	}
	return nil
}
`

const BidirectionalStreamingRPCWithErrorsServerInterfaceCode = `// MethodBidirectionalStreamingRPCWithErrors implements the
// "MethodBidirectionalStreamingRPCWithErrors" method in
// pb.ServiceBidirectionalStreamingRPCWithErrorsServer interface.
func (s *Server) MethodBidirectionalStreamingRPCWithErrors(stream pb.ServiceBidirectionalStreamingRPCWithErrors_MethodBidirectionalStreamingRPCWithErrorsServer) error {
	ctx := stream.Context()
	ctx = context.WithValue(ctx, goa.MethodKey, "MethodBidirectionalStreamingRPCWithErrors")
	ctx = context.WithValue(ctx, goa.ServiceKey, "ServiceBidirectionalStreamingRPCWithErrors")
	p, err := s.MethodBidirectionalStreamingRPCWithErrorsH.Decode(ctx, nil)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "timeout":
				return goagrpc.NewStatusError(codes.Canceled, err)
			case "internal":
				return goagrpc.NewStatusError(codes.Unknown, err)
			case "bad_request":
				return goagrpc.NewStatusError(codes.InvalidArgument, err)
			}
		}
		return goagrpc.EncodeError(err)
	}
	ep := &servicebidirectionalstreamingrpcwitherrors.MethodBidirectionalStreamingRPCWithErrorsEndpointInput{
		Stream: &MethodBidirectionalStreamingRPCWithErrorsServerStream{stream: stream},
	}
	err = s.MethodBidirectionalStreamingRPCWithErrorsH.Handle(ctx, ep)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "timeout":
				return goagrpc.NewStatusError(codes.Canceled, err)
			case "internal":
				return goagrpc.NewStatusError(codes.Unknown, err)
			case "bad_request":
				return goagrpc.NewStatusError(codes.InvalidArgument, err)
			}
		}
		return goagrpc.EncodeError(err)
	}
	return nil
}
`