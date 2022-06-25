export interface Position {
    latitude: number;
    longitude: number;
}

export interface Route {
    id: string;
    title: string;
    startPosition: Position;
    endPosition: Position
}