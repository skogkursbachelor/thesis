const [mapInstance, setMapInstance] = useState<OpenLayersMap | null>(null);

[...]

const map = new OpenLayersMap({
    target: mapRef.current,
    view: new View({
      center: fromLonLat([10, 59]),
      zoom: 5,
    }),
    controls: [attribution],
    layers: layers,
  });

  setMapInstance(map);

[...]

<BaseLayerGroup map={mapInstance} />
<MapZoom map={mapInstance} />
