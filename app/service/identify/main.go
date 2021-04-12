package identify
//
//// GetTokenInfo verify user info by accesskey.
//func (s *Service) GetTokenInfo(c context.Context, token *v1.GetTokenInfoReq) (res *model.IdentifyInfo, err error) {
//	var cache = true
//	if res, err = s.d.AccessCache(c, token.Token); err != nil {
//		cache = false
//	}
//	if res != nil {
//		if res.Mid == _noLoginMid {
//			err = ecode.NoLogin
//			return
//		}
//		s.loginLog(res.Mid, metadata.String(c, metadata.RemoteIP), metadata.String(c, metadata.RemotePort), token.Buvid)
//		return
//	}
//	if res, err = s.d.AccessToken(c, token.Token); err != nil {
//		if err != ecode.NoLogin && err != ecode.AccessKeyErr {
//			return
//		}
//		// no login and need cache 30s
//		res = _noLoginIdentify
//	}
//	if cache && res != nil {
//		s.cache.Save(func() {
//			s.d.SetAccessCache(context.Background(), token.Token, res)
//		})
//		// if cache err or res nil, don't call addLoginLog
//		s.loginLog(res.Mid, metadata.String(c, metadata.RemoteIP), metadata.String(c, metadata.RemotePort), token.Buvid)
//	}
//	if res.Mid == _noLoginMid {
//		err = ecode.NoLogin
//		return
//	}
//	return
//}
