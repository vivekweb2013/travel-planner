import { Autocomplete, FormControl, Grid, TextField } from '@mui/material';
import 'mapbox-gl/dist/mapbox-gl.css';
import React from 'react';
import ReactMapboxGl, { Feature, Layer } from 'react-mapbox-gl';
import { useAppDispatch, useAppSelector } from '../../app/hooks';
import {
  City,
  fetchCitiesAsync, fetchTravelPlanAsync, selectCities, selectRoutes
} from './travelPlanSlice';

const Map = ReactMapboxGl({
  accessToken:
    "pk.eyJ1IjoibWF5b2ppY2giLCJhIjoiY2pla3Q3MzVvMWRoYTJybnVyMndxM2hmdCJ9.nWZlYcpKaMqz6m7xTsnJGA"
});

const lineLayout = {
  "line-join": "round",
  "line-cap": "round"
}
const linePaint = {
  "line-color": "#ff11ff",
  "line-width": 4,
  "line-opacity": 1
}

const getCirclePaint = () => ({
  'circle-radius': 30,
  'circle-color': '#E54E52',
  'circle-opacity': 0.8
});

export function TravelPlan() {
  const cities = useAppSelector(selectCities);
  const routes = useAppSelector(selectRoutes);
  const mappedRoutes = routes.map(r => [r.source.location.lon, r.source.location.lat] as [number, number])
  console.log(mappedRoutes)
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
          isOptionEqualToValue={(o, v) => o.name === v.name}
          getOptionLabel={(option) => option.name}
          renderOption={(props, option) => <li {...props}>{option.name}</li>}
          renderInput={(params) => <TextField {...params} label="Select City" />} value={city}
          onChange={(event: any, newValue: City | null) => { handleValueChange(newValue); }}
          onInputChange={(e, v) => handleQueryChange(v)}
        />
      </FormControl>

      {routes.length > 0 && <Grid container spacing={2}>
        <Grid item xs={12}>
          {routes.map(r => <span key={r.source.name}>{r.source.name} {"->"} </span>)} {routes[routes.length - 1].dest.name}
        </Grid>
        <Grid item xs={12}>
          Total Distance: {routes.map(r => r.distance).reduce((s, a) => s + a, 0).toFixed(2)} km
        </Grid>
        <Grid container spacing={0} direction="column" alignItems="center" justifyContent="center" style={{ height: '100vh' }}>
          <Map style="mapbox://styles/mapbox/streets-v9" containerStyle={{ height: '100%', width: '80%' }} zoom={[0]}>
            <Layer type="line" layout={lineLayout} paint={linePaint}>
              <Feature coordinates={[[-0.127758, 51.507351], [72.877655, 19.075983]]} />
            </Layer>
          </Map>
        </Grid>
      </Grid>
      }

    </div>
  );
}
