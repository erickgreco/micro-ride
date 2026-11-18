export const API_URL = process.env.NEXT_PUBLIC_API_URL ?? 'http://localhost:8081';
export const WEBSOCKET_URL = process.env.NEXT_PUBLIC_WEBSOCKET_URL ?? 'ws://localhost:8081/ws';

export const CARTO_API_KEY = process.env.NEXT_PUBLIC_CARTO_API_KEY ?? '';
export const MAP_TILES_URL = `https://basemaps.cartocdn.com/rastertiles/voyager/{z}/{x}/{y}.png?key=${CARTO_API_KEY}`;
