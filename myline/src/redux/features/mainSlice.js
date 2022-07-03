import { createSlice, createAsyncThunk } from '@reduxjs/toolkit'
import { GET, POST } from '../../utils/ApiClient'

const initialState = {
    main: null,
    mainStatus: { status: "idle", error: "" },
}



const mainSlice = createSlice({
    name: "main",
    initialState,
    reducers: {
        setLineToken(state, action) {
           
        },
    },
    extraReducers: {
       

    }
})

export const selectedMain = (state) => state.mainSlice;

export const {
    setLineToken
} = mainSlice.actions

export default mainSlice.reducer