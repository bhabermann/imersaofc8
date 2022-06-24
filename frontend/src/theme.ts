import { createTheme, PaletteOptions } from "@mui/material";
import { grey, yellow } from "@mui/material/colors";

const palette: PaletteOptions = {
    mode: "dark",
    primary: {
        main: yellow[700],//"#FFCD00",
        contrastText: grey[900] // "#232526",
    },
    background: {
        default: grey[900] //"#242526"
    },
}

const theme = createTheme({
    palette,
});

export default theme;