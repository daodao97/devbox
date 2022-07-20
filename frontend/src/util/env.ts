export function isDevMode() {
  return import.meta.env.DEV
}

export function isProdMode() {
  return import.meta.env.PROD
}
