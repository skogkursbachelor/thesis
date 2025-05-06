for i := 0; i < sf.NumRecords(); i++ {
    atr, geometry := sf.Record(i)
    switch g := geometry.(type) {
    case *geom.Polygon:
        bbox := g.Bounds()
        k := fmt.Sprintf("%s_%d", f, i)
        index.Insert(bbox.Min(0), bbox.Min(1), bbox.Max(0), bbox.Max(1), k, atr)
    case *geom.MultiPolygon:
        for j := 0; j < g.NumPolygons(); j++ {
            polygon := g.Polygon(j)
            bbox := polygon.Bounds()
            k := fmt.Sprintf("%s_%d_%d", f, i, j)
            index.Insert(bbox.Min(0), bbox.Min(1), bbox.Max(0), bbox.Max(1), k, atr)
        }
    case *geom.MultiLineString:
        for j := 0; j < g.NumLineStrings(); j++ {
            line := g.LineString(j)
            bbox := line.Bounds()
            k := fmt.Sprintf("%s_%d_%d", f, i, j)
            index.Insert(bbox.Min(0), bbox.Min(1), bbox.Max(0), bbox.Max(1), k, atr)
        }
    default:
        log.Debug().Msgf("Unsupported geometry type in %s: %T", f, geometry)
    }
}