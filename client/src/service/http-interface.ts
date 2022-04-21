import { HttpService } from './http-service';

export interface HttpInterface {
    HttpGetRequest(url: string): Promise<unknown>;
}

export const NewHttpInterface = (): HttpInterface | undefined => {
    return new HttpService();
}