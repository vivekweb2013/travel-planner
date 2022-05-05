import { City } from "./travelPlanSlice";

const API_URL = "/api/v1";

export const getMatchingCities = (query: string): Promise<City[]> => {
  return fetch(`${API_URL}/cities?query=${query}`).then(async (res) => {
    if (!res.ok) {
      return Promise.reject(await res.json());
    }
    return await res.json();
  })
}