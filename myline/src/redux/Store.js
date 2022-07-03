import {
    configureStore,
    applyMiddleware,
    combineReducers,
    createStore,
} from '@reduxjs/toolkit';

import mainReducer from "./features/mainSlice";

const reducers = combineReducers({
    main: mainReducer,
});

const store = configureStore({
    reducer: reducers,
});

export default store;