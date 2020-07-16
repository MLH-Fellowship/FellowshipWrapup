import styled from "styled-components";

export default function GithubProfile({ info }) {
  console.log(info);

  return (
    <StyledLink href="https://github.com/sebastiancrossa" target="_blank">
      <StyledBackground>
        <div className="image-container">
          <img src="/gh-logo.png" alt="Github logo" />
        </div>
        <div className="text-container">
          <p className="head">Sebastian Crossa | Zapopan, Jalisco</p>
          <p>Computer science student & maker - Full-stack JavaScript</p>
        </div>
      </StyledBackground>
    </StyledLink>
  );
}

export const StyledBackground = styled.div`
  display: flex;
  align-items: center;

  border: 1px solid gray;
  border-radius: 8px;

  padding: 1rem;
  max-width: 20.5rem;

  color: gray;
  background-color: var(--color-gray);

  p {
    margin: 0;
  }

  .head {
    color: var(--color-text);
    font-weight: 600;
  }

  .image-container {
    background-color: var(--color-gray);
    margin-right: 0.5rem;

    img {
      width: 2rem;
      height: auto;
    }
  }
`;

export const StyledLink = styled.a`
  &:hover {
    text-decoration: none;
  }
`;

export async function getServerSideProps() {
  // ! Getting a weid error when calling the below func, needs fixing
  //const info = await getUserInfo();

  return {
    props: {
      info,
    },
  };
}
