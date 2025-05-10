/* ./src/utils/trafficability.ts */
function getTrafficabilityColor(
  frostDepth: number,
  frostDepthThreshold = 10,
  waterSaturation: number,
  minSaturationThreshold = 65,
  maxSaturationThreshold = 80
): number[] {
  if (frostDepth >= frostDepthThreshold) {
    return green;
  } else if (waterSaturation >= maxSaturationThreshold) {
    return red;
  } else if (waterSaturation >= minSaturationThreshold) {
    return yellow;
  } else {
    return green;
  }
}