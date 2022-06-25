import './App.css';
import { Mapping } from './components/Mapping';
import { ThemeProvider } from '@emotion/react';
import theme from './theme';
import { CssBaseline } from '@mui/material';

function App() {
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Mapping />
    </ThemeProvider>
      
    // <div className="App">
    //   <header className="App-header">
    //     <img src={logo} className="App-logo" alt="logo" />
    //     <p>
    //       Edit <code>src/App.tsx</code> and save to reload.
    //     </p>
    //     <a
    //       className="App-link"
    //       href="https://reactjs.org"
    //       target="_blank"
    //       rel="noopener noreferrer"
    //     >
    //       Learn React with FC8
    //     </a>
    //   </header>
    // </div>
  );
}

export default App;
