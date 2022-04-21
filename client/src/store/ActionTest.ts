import {Action, Reducer} from "redux";
import {AppThunkAction} from "./index";
import axios, {Method} from "axios";

export interface ActionTestState {
  isLoading: boolean;
  reqMethod: string;
  reqPath: string;
  reqHeaders?: { [key: string]: string };
  reqBody?: Record<string, unknown>;
  resBody?: Record<string, unknown>;
}

interface RequestActionTestAction {
  type: 'REQUEST_ACTION_TEST';
  reqMethod: string;
  reqPath: string;
  reqHeaders?: { [key: string]: string };
  reqBody?: Record<string, unknown>;
}

interface ReceiveActionTestAction {
  type: 'RECEIVE_ACTION_TEST';
  resBody: Record<string, unknown>;
}

type KnownAction = RequestActionTestAction | ReceiveActionTestAction;

export const actionCreators = {
  requestActionTest: (reqMethod: string,
    reqPath: string,
    reqHeaders?: { [key: string]: string },
    reqBody?: Record<string, unknown>)
    : AppThunkAction<KnownAction> => (dispatch, getState) => {
    const appState = getState();
    if (appState && appState.actionTest && reqPath !== appState.actionTest.reqPath) {
      axios(reqPath, {
        method: reqMethod as Method,
        headers: reqHeaders,
        data: reqBody === undefined ? undefined : JSON.stringify(reqBody)
      }).then(response => {
        dispatch({type: 'RECEIVE_ACTION_TEST', resBody: response.data as Record<string, unknown>});
      }).catch(error => {
        console.log(error?.toJSON());
      })

      dispatch({
        type: 'REQUEST_ACTION_TEST',
        reqMethod: reqMethod,
        reqPath: reqPath,
        reqHeaders: reqHeaders,
        reqBody: reqBody
      });
    }
  }
}

const unloadedState: ActionTestState = {
  isLoading: false,
  reqMethod: '',
  reqPath: '',
  reqHeaders: undefined,
  reqBody: undefined,
  resBody: undefined
}

export const reducer: Reducer<ActionTestState> = (state: ActionTestState | undefined, incomingAction: Action)
  : ActionTestState => {
  if (state === undefined) {
    return unloadedState;
  }

  const action = incomingAction as KnownAction;
  switch (action.type) {
  case "REQUEST_ACTION_TEST":
    return {
      isLoading: true,
      reqMethod: action.reqMethod,
      reqPath: action.reqPath,
      reqHeaders: action.reqHeaders,
      reqBody: action.reqBody
    }

  case "RECEIVE_ACTION_TEST":
    return {
      isLoading: false,
      reqMethod: state.reqMethod,
      reqPath: state.reqPath,
      reqHeaders: state.reqHeaders,
      reqBody: state.reqBody,
      resBody: action.resBody,
    }
  }

  return state;
}
