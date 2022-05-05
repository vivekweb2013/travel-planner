import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import { RootState } from '../../app/store';
import { getMatchingCities, getTravelPlan } from './travelPlanAPI';


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

export interface Route {
  order: number
  source: City
  dest: City
  distance: number
}

export interface TravelPlanState {
  cities: City[],
  routes: Route[],
  status: 'idle' | 'loading' | 'failed';
}

const initialState: TravelPlanState = {
  cities: [],
  routes: [],
  status: 'idle',
};

export const fetchCitiesAsync = createAsyncThunk(
  'travelPlan/fetchCities',
  async (query: string) => {
    const response = await getMatchingCities(query);
    return response;
  }
);

export const fetchTravelPlanAsync = createAsyncThunk(
  'travelPlan/fetchTravelPlan',
  async (city: City) => {
    const response = await getTravelPlan(city);
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
        state.cities = initialState.cities;
      })
      .addCase(fetchCitiesAsync.fulfilled, (state, action) => {
        state.status = 'idle';
        state.cities = action.payload;
      })
      .addCase(fetchCitiesAsync.rejected, (state) => {
        state.status = 'failed';
        state.cities = initialState.cities;
      })

      .addCase(fetchTravelPlanAsync.pending, (state) => {
        state.status = 'loading';
        state.routes = initialState.routes;
      })
      .addCase(fetchTravelPlanAsync.fulfilled, (state, action) => {
        state.status = 'idle';
        state.routes = action.payload;
      })
      .addCase(fetchTravelPlanAsync.rejected, (state) => {
        state.status = 'failed';
        state.routes = initialState.routes;
      });
  },
});

export const selectCities = (state: RootState) => state.travelPlan.cities;
export const selectRoutes = (state: RootState) => state.travelPlan.routes;

export default travelPlanSlice.reducer;
