import { createTheme, PaletteOptions } from "@mui/material";

const palette: PaletteOptions = {
    // type: "dark",
    primary: {
        main: "#FFCD00",
        contrastText: "#232526",
    },
    background: {
        default: "#242526"
    },
}

const theme = createTheme({
    palette,
});

export default theme;