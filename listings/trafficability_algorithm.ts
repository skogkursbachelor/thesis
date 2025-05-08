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
  }
  if (waterSaturation >= maxSaturationThreshold) {
    return red;
  }
  if (waterSaturation >= minSaturationThreshold) {
    return yellow;
  }
  return green;
}