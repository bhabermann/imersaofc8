import { Button, Grid, MenuItem, Select } from "@mui/material"

type Props = {}

export const Mapping = (props: Props) => {
    return (
        <Grid container>
            <Grid item xs={12} sm={3}>
                <form>
                    <Select fullWidth>
                        <MenuItem value="">
                            <em>Selecione uma corrida</em>
                        </MenuItem>
                        <MenuItem value="1">
                            <em>Primeira corrida</em>
                        </MenuItem>
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