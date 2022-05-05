import { Autocomplete, FormControl, TextField } from '@mui/material';
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

  const handleQueryChange = (newValue: City | null) => {
    setCity(newValue);
  }
  useEffect(() => {
    dispatch(fetchCitiesAsync("New"));
  }, [dispatch])

  return (
    <div>
      <FormControl sx={{ m: 1, width: 300 }}>
        <Autocomplete id="city-select" options={cities}
          getOptionLabel={(option) => option.name}
          renderOption={(props, option) => <li {...props}>{option.name}</li>}
          renderInput={(params) => <TextField {...params} label="Select City" />} value={city}
          onChange={(event: any, newValue: City | null) => { handleQueryChange(newValue); }}
        />
      </FormControl>
    </div>
  );
}
