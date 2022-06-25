import { Button, Grid, MenuItem, Select } from "@mui/material"
import { useEffect, useState } from "react"
import { Route } from "../util/models";

const API_URL = process.env.REACT_APP_API_URL

type Props = {}

export const Mapping = (props: Props) => {
    const [routes, setRoutes] = useState<Route[]>([]);

    useEffect(() => {
        fetch(`${API_URL}/routes`)
            .then((data) => data.json())
            .then((data) => setRoutes(data));
    }, [])

    return (
        <Grid container>
            <Grid item xs={12} sm={3}>
                <form>
                    <Select fullWidth>
                        <MenuItem value="">
                            <em>Selecione uma corrida</em>
                        </MenuItem>
                        { routes.map((route, key) => (
                            <MenuItem key={key} value={route.id}>
                                <em>{route.title}</em>
                            </MenuItem>
                        ))}

                        <MenuItem value="2">
                            <em>Segunda corrida</em>
                        </MenuItem>
                        <MenuItem value="3">
                            <em>Terceira corrida</em>
                        </MenuItem>                                                                        
                    </Select>
                    <Button type="submit" color="primary" variant="contained">
                        Iniciar uma corrida
                    </Button>
                </form>
            </Grid>
            <Grid item xs={12} sm={9}>
                <div id="map"></div>
            </Grid>
        </Grid>
    )
}