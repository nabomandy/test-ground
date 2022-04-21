import { HttpInterface } from './http-interface';
import axios, { AxiosError } from 'axios';

const HTTP_REQUEST_TIMEOUT = 12000; // HTTP timeout 시간 설정

export class HttpService implements HttpInterface {
    /**
     * @method HttpGetRequest
     * @description HTTP GET 요청 메소드
     * @returns 
     */
    public async HttpGetRequest(url: string): Promise<unknown> {
        const response = await axios(url, {
            method: 'GET',
            headers: {},
            timeout: HTTP_REQUEST_TIMEOUT,
        }).catch((err: Error | AxiosError) => {
            if (axios.isAxiosError(err))  {
                console.log(`Method: HttpGetRequest, url: ${url} => ${err}`);
            }
        });

        if (response && response.status == 200) {
            return response.data;
        }

        return Promise.resolve(undefined);
    }
}