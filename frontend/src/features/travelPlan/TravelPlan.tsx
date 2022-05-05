import { Autocomplete, FormControl, InputLabel, TextField } from '@mui/material';
import React, { useEffect } from 'react';
import { useAppDispatch, useAppSelector } from '../../app/hooks';
import {
  City,
  fetchCitiesAsync, selectCities
} from './travelPlanSlice';


export function TravelPlan() {
  const cities = useAppSelector(selectCities);
  const dispatch = useAppDispatch();
  const [city, setCity] = React.useState<City | null>();

  useEffect(() => {
    dispatch(fetchCitiesAsync(""));
  }, [])

  return (
    <div>
      <FormControl sx={{ m: 1, width: 300 }}>
        <InputLabel id="demo-multiple-name-label">Name</InputLabel>
        <Autocomplete
          id="city-select"
          options={cities}
          renderInput={(params) => <TextField {...params} label="Select City" />}
          value={city}
          onChange={(event: any, newValue: City | null) => {
            setCity(newValue);
          }}
        />
      </FormControl>
    </div>
  );
}
