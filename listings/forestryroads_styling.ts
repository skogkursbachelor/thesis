/* ./src/components/map/layers/ForestryRoadLayer.tsx */
const roadStyle = (feature: FeatureLike) => {
  const frostDepth = feature.get("teledybde");
  const soilSaturation = feature.get("vannmetning");
  const depositType = feature.get("løsmassekoder");
  const thresholds = currentThresholds.get(depositType[0]);
  const minThreshold = thresholds?.min ?? 65;
  const maxThreshold = thresholds?.max ?? 80;
  const frostDepthThreshold = currentFrostDepthThreshold;

  const color = getTrafficabilityColor(
    frostDepth,
    frostDepthThreshold,
    soilSaturation,
    minThreshold,
    maxThreshold
  );

  // Increase width on hover
  const width = feature === hoveredFeature ? 6 : 3;

  return [
    new Style({
      stroke: new Stroke({
        color: color,
        width: width,
      }),
    }),
  ];
};