export interface QueryParams {
  page: number;
  limit: number;
}

export interface Paginated<T> {
  metadata: {
    size: number;
    page: number;
    hasNext: boolean;
    total: number;
  };
  data: T[];
}
