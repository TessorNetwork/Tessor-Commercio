package cli_test

// func networkWithDocumentObjects(t *testing.T, n int) (*network.Network, []*types.DidDocument) {
// 	t.Helper()
// 	cfg := network.DefaultConfig()
// 	state := types.GenesisState{}
// 	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

// 	for i := 0; i < n; i++ {
// 		_, _, addr := testdata.KeyTestPubAddr()
// 		pubkeys := []*types.PubKey{}
// 		service := []*types.Service{}
// 		state.DidDocuments = append(state.DidDocuments, &types.DidDocument{ID: addr.String(), PubKeys: pubkeys, Service: service})
// 	}
// 	buf, err := cfg.Codec.MarshalJSON(&state)
// 	require.NoError(t, err)

// 	stateGov := govTypes.GenesisState{}
// 	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[govTypes.ModuleName], &stateGov))
// 	stateGov.GovernmentAddress = "cosmos1nynns8ex9fq6sjjfj8k79ymkdz4sqth06xexae"
// 	bufGov, err := cfg.Codec.MarshalJSON(&stateGov)
// 	require.NoError(t, err)

// 	cfg.GenesisState[types.ModuleName] = buf
// 	cfg.GenesisState[govTypes.ModuleName] = bufGov

// 	return network.New(t, cfg), state.DidDocuments
// }

// func TestShowIdentity(t *testing.T) {
// 	net, objs := networkWithDocumentObjects(t, 2)

// 	ctx := net.Validators[0].ClientCtx
// 	common := []string{
// 		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
// 	}
// 	for _, tc := range []struct {
// 		desc string
// 		id   string
// 		args []string
// 		err  error
// 		obj  *types.DidDocument
// 	}{
// 		{
// 			desc: "found",
// 			id:   fmt.Sprintf("%s", objs[0].ID),
// 			args: common,
// 			obj:  objs[0],
// 		},
// 		{
// 			desc: "not found",
// 			id:   "not_found",
// 			args: common,
// 			err:  status.Error(codes.InvalidArgument, "not found"),
// 		},
// 	} {
// 		tc := tc
// 		t.Run(tc.desc, func(t *testing.T) {
// 			args := []string{tc.id}
// 			args = append(args, tc.args...)
// 			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowIdentity(), args)
// 			if tc.err != nil {
// 				stat, ok := status.FromError(tc.err)
// 				require.True(t, ok)
// 				require.ErrorIs(t, stat.Err(), tc.err)
// 			} else {
// 				require.NoError(t, err)
// 				var resp types.QueryResolveDidResponse
// 				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
// 				require.NotNil(t, resp.DidDocument)
// 				require.Equal(t, tc.obj, resp.DidDocument)
// 			}
// 		})
// 	}
// }
