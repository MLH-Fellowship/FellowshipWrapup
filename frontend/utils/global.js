import { createGlobalStyle } from "styled-components";

export default createGlobalStyle`
    *,
    *::before,
    *::after {
        margin: 0;
        padding: 0;
        box-sizing: border-box;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
    }

    html,
    body {
        padding: 0;
        margin: 0;
        font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Roboto, Oxygen, Ubuntu, Cantarell, Fira Sans, Droid Sans, Helvetica Neue, sans-serif;
        font-weight: 400;

        a {
            color: ${({ theme }) => theme.color.main};
            font-weight: 600;
        }

        code {
          background: #fafafa;
          border-radius: 5px;
          padding: 0.75rem;
          font-size: 1.1rem;
          font-family: Menlo, Monaco, Lucida Console, Liberation Mono,
            DejaVu Sans Mono, Bitstream Vera Sans Mono, Courier New, monospace;
        }
    }

    html {
        --color-text: ${({ theme }) => theme.color.text};
        --color-white: ${({ theme }) => theme.color.white};
        --color-gray: ${({ theme }) => theme.color.gray};
        --color-black: ${({ theme }) => theme.color.black};
        --color-main: ${({ theme }) => theme.color.main};
        --color-secondary: ${({ theme }) => theme.color.secondary};
    }
`;
