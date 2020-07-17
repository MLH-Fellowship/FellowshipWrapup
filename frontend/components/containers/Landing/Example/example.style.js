import styled from "styled-components";

export const StyledBackground = styled.section`
  background: var(--color-main) url(${(props) => props.img}) center / cover;

  margin: 0 auto;
  padding: 2.5rem;

  max-width: 95%;
  min-height: 45vh;

  -webkit-box-shadow: 0px 1px 80px -16px rgba(130, 130, 130, 0.5);
  -moz-box-shadow: 0px 1px 80px -16px rgba(130, 130, 130, 0.5);
  box-shadow: 0px 1px 80px -16px rgba(130, 130, 130, 0.5);

  border-radius: 8px;
`;
