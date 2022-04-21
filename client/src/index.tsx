import * as React from 'react';
import * as ReactDOM from 'react-dom';
import { Provider} from 'react-redux';
import { ConnectedRouter } from 'connected-react-router';
import { createBrowserHistory } from 'history';
import configureStore from './store/configureStore';
import App from './App';
import registerServiceWorker from './registerServiceWorker';
import moment from 'moment';
import 'moment/locale/ko'

import './index.scss'

moment.locale('ko');

// Redux Store에서 사용할 브라우저 history 생성
const baseUrl = document.getElementsByTagName('base')[0].getAttribute('href') as string;
const history = createBrowserHistory({basename: baseUrl});

// Get the application-wide store instance, pre-populating with state from the server where available.
const store = configureStore(history);

ReactDOM.render(
	<Provider store={store}>
		<ConnectedRouter history={history}>
			<App/>
		</ConnectedRouter>
	</Provider>,
	document.getElementById('root')
);

registerServiceWorker();
