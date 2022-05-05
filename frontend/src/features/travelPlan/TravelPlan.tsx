import { Autocomplete, FormControl, Grid, TextField } from '@mui/material';
import React from 'react';
import { useAppDispatch, useAppSelector } from '../../app/hooks';
import {
  City,
  fetchCitiesAsync, fetchTravelPlanAsync, selectCities, selectRoutes
} from './travelPlanSlice';


export function TravelPlan() {
  const cities = useAppSelector(selectCities);
  const routes = useAppSelector(selectRoutes);
  const dispatch = useAppDispatch();
  const [city, setCity] = React.useState<City | null>();

  const handleValueChange = (city: City | null) => {
    setCity(city);
    city && dispatch(fetchTravelPlanAsync(city));
  }

  const handleQueryChange = (value: string) => {
    value && dispatch(fetchCitiesAsync(value));
  }

  return (
    <div>
      <FormControl sx={{ m: 1, width: 300 }}>
        <Autocomplete id="city-select" options={cities}
          getOptionLabel={(option) => option.name}
          renderOption={(props, option) => <li {...props}>{option.name}</li>}
          renderInput={(params) => <TextField {...params} label="Select City" />} value={city}
          onChange={(event: any, newValue: City | null) => { handleValueChange(newValue); }}
          onInputChange={(e, v) => handleQueryChange(v)}
        />
      </FormControl>

      {routes.length > 0 && <Grid container spacing={2}>
        <Grid item xs={12}>
          {routes.map(r => <span>{r.source.name} {"->"} </span>)} {routes[routes.length - 1].dest.name}
        </Grid>
        <Grid item xs={12}>
          Total Distance: {routes.map(r => r.distance).reduce((s, a) => s + a, 0).toFixed(2)} km
        </Grid>
        <Grid item xs={12}>

        </Grid>
      </Grid>
      }
    </div>
  );
}
