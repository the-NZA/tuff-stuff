export type Response<T> = {
	data: T;
	code: number;
	error: string;
};
