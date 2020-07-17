import styled from "styled-components";

export default function Map() {
  return (
    <StyledBackground>
      <div className="text-container">
        <h1>
          With over <span>140+ students</span> from over <span>23 regions</span>{" "}
          and <span>15 different timezones</span>, I was able to network and
          reach out to fellows from all around the world!
        </h1>
      </div>

      <div className="image-container">
        <img src="/gh-fellow-transparent.png" alt="Pixelated fellowship map" />
      </div>
    </StyledBackground>
  );
}

const StyledBackground = styled.div`
  padding: 1.5rem;
  text-align: center;
  margin-bottom: 13rem;
  background-color: #f5f5f9;
  border-radius: 8px;

  h1 {
    font-size: 2rem;
    margin-bottom: 1.5rem;
  }

  span {
    color: var(--color-green-light);
    font-weight: 600;
  }

  @media (max-width: 412px) {
    h1 {
      font-size: 1.2rem;
    }
  }

  .text-container {
    max-width: 50rem;
    margin: 0 auto;
  }

  .image-container {
    padding: 1rem;
    max-width: 50rem;
    margin: 0 auto;

    border-radius: 8px;
    background-color: var(--color-gray);

    img {
      max-width: 100%;
    }
  }
`;
