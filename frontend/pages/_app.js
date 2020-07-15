// --- Styled Components -- //
import { ThemeProvider } from "styled-components";
import theme from "../utils/theme";
import GlobalStyles from "../utils/global";
import "../styles/userStyled.css";

export default function App({ Component, pageProps }) {
  return (
    <ThemeProvider theme={theme}>
      <>
        <Component {...pageProps} />
        <GlobalStyles />
      </>
    </ThemeProvider>
  );
}
