import styled from "styled-components";

export const StyledHeader = styled.header`
  text-align: left;
  margin-bottom: 7rem;

  img {
    height: 22rem;
    border-radius: 8px;
    margin: 0 auto 1rem auto;

    -webkit-box-shadow: 0px 1px 80px -16px rgba(130, 130, 130, 1);
    -moz-box-shadow: 0px 1px 80px -16px rgba(130, 130, 130, 1);
    box-shadow: 0px 1px 80px -16px rgba(130, 130, 130, 1);
  }

  .textContainer {
    max-width: 60rem;
    margin: 0 auto;
    padding: 0 3rem;
    color: var(--color-text);

    h1 {
      font-weight: 600;
      font-size: 1.7rem;
      margin-bottom: 0.5rem;
    }

    h2 {
      font-weight: 600;
      font-size: 2rem;
    }

    a {
      text-decoration: underline;
    }

    a,
    span {
      color: var(--color-main);
    }
  }
`;

export const Grid = styled.div`
  display: grid;
  grid-template-columns: 60% 40%;
  align-items: center;

  padding: 1.5rem 0 1.5rem 0;
`;
