export const getPreviousTimeSpan = (
  timestamp: number,
  timeSpan: number
): number => {
  timestamp
  return Math.ceil(timestamp / timeSpan) * timeSpan - timeSpan
}
