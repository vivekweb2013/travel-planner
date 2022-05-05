import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import { RootState } from '../../app/store';
import { getMatchingCities } from './travelPlanAPI';


export interface City {
  id: string
  name: string
  countryName: string
  contId: number
  location: {
    lat: number,
    lon: number
  }
}

export interface TravelPlanState {
  cities: City[],
  status: 'idle' | 'loading' | 'failed';
}

const initialState: TravelPlanState = {
  cities: [],
  status: 'idle',
};

export const fetchCitiesAsync = createAsyncThunk(
  'travelPlan/fetchCities',
  async (query: string) => {
    const response = await getMatchingCities(query);
    return response;
  }
);

export const travelPlanSlice = createSlice({
  name: 'travelPlan',
  initialState,
  reducers: {
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchCitiesAsync.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(fetchCitiesAsync.fulfilled, (state, action) => {
        state.status = 'idle';
        state.cities = action.payload;
      })
      .addCase(fetchCitiesAsync.rejected, (state) => {
        state.status = 'failed';
      });
  },
});

export const selectCities = (state: RootState) => state.travelPlan.cities;

export default travelPlanSlice.reducer;
