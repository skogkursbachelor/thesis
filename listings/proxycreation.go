mux := http.NewServeMux()

proxies, err := utils.LoadProxiesFromFile()
if err != nil {
    log.Fatal().Msg("Error loading proxies: " + err.Error())
}

for path, remoteAddr := range proxies {
    log.Info().Msg(path + "->" + remoteAddr)
    p := &handlers.Proxy{RemoteAddr: remoteAddr}
    mux.HandleFunc(constants.ProxyPath+path, p.ProxyHandler)
}

mux.HandleFunc(constants.BaseLayerPath+"/{type}/{abc}/{z}/{x}/{y}",
    handlers.BaseLayerHandler)
mux.HandleFunc(constants.BaseLayerPath+"/{type}/{abc}/{z}/{x}", 
    handlers.BaseLayerHandler)