import styled from "styled-components";

const GithubProfile = ({ accountInfo }) => {
  const { Name, Location, Bio } = accountInfo;

  return (
    <StyledLink href="https://github.com/sebastiancrossa" target="_blank">
      <StyledBackground>
        <div className="image-container">
          <img src="/gh-logo.png" alt="Github logo" />
        </div>
        <div className="text-container">
          <p className="head">
            {Name} {Location && `| ${Location}`}
          </p>
          <p>{Bio}</p>
        </div>
      </StyledBackground>
    </StyledLink>
  );
};

export const StyledBackground = styled.div`
  display: flex;
  align-items: center;
  text-align: left;

  border: 1px solid gray;
  border-radius: 8px;

  padding: 1rem;

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

export default GithubProfile;
