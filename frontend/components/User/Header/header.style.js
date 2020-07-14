import styled from "styled-components";

export const StyledBackground = styled.div`
  text-align: left;
  padding: 2.5rem 4rem;

  max-width: 95%;
  margin: 2rem auto;

  border-radius: 8px;

  background-color: var(--color-gray);

  .textContainer {
    display: flex;

    span {
      color: var(--color-main);
    }

    img {
      height: 7rem;
      border-radius: 8px;

      margin-right: 1rem;
    }
  }

  .underline {
    text-decoration: underline;
  }
`;
