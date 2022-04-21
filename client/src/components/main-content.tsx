import React, { useState } from 'react';
import { NewHttpInterface } from '../service/http-interface';

import './main-content.scss';

interface httpResultInterface {
	code: number;
	message: string;
	response: unknown;
	responseTime: string;
}

export const MainContent = (): JSX.Element => {
	const [ result, setResult ] = useState('');

	const getHttpRequest = async (url: string) => {
		const httpService = NewHttpInterface();
		const res = await httpService.HttpGetRequest(url) as httpResultInterface;
		if (res) {
			setResult(JSON.stringify(res.response, null, 4));
		}
	};

	return (
		<div className="main-content">
			<div className="result-btn-group">
				<button onClick={() => getHttpRequest('/vurix-dms/api/v1/dbData/getDataList')}>[GET] TEST</button>
				<button onClick={() => getHttpRequest('/vurix-dms/api/v1/apiTest/getAPIData?name=AAAA')}>[GET] API1</button>
				<button onClick={() => getHttpRequest('/vurix-dms/api/v1/apiTest/getAPIData2?name=AAAA')}>[GET] API2</button>
			</div>
			{ 
				result 
					? <div className="result-box"><pre>{result}</pre></div>
					: <div className="no-result-box"><img alt="noResult" src='assets/images/no_results.jpg'></img></div>
			}
		</div>
	)
}