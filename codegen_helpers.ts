export function capitalize(str: string): string {
  if (typeof str !== "string") return "";
  if (str.length === 0) return "";
  return str.charAt(0).toUpperCase() + str.slice(1);
}
