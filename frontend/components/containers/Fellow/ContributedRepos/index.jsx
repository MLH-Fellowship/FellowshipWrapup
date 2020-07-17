import styled from "styled-components";

export const ContributedRepos = ({ repoContribs }) => {
//   const { Name, Url } = repoContribs;


  return (
    <StyledRepo>
      <a>
        {repoContribs.nodes.map((repo) => {
          return (
            <div className="repo-card-div">
              <div className="repo-name-div">
                <svg
                  aria-hidden="true"
                  className="octicon"
                  height="20"
                  role="img"
                  viewBox="0 0 12 16"
                  width="14"
                  className="repo-svg"
                >
                  <path
                    fill-rule="evenodd"
                    d="M4 9H3V8h1v1zm0-3H3v1h1V6zm0-2H3v1h1V4zm0-2H3v1h1V2zm8-1v12c0 .55-.45 1-1 1H6v2l-1.5-1.5L3 16v-2H1c-.55 0-1-.45-1-1V1c0-.55.45-1 1-1h10c.55 0 1 .45 1 1zm-1 10H1v2h2v-1h3v1h5v-2zm0-10H2v9h9V1z"
                  ></path>
                </svg>
                <p className="repo-name">{repo.name}</p>
              </div>
              <p className="repo-description">{repo.description}</p>
              <div className="repo-stats">
                <div className="repo-left-stat">
                  <span>
                    <svg
                      aria-hidden="true"
                      className="octicon"
                      height="20"
                      role="img"
                      viewBox="0 0 10 16"
                      width="12"
                      fill="rgb(106, 115, 125)"
                      className="repo-star-svg"
                    >
                      <path
                        fill-rule="evenodd"
                        d="M8 1a1.993 1.993 0 0 0-1 3.72V6L5 8 3 6V4.72A1.993 1.993 0 0 0 2 1a1.993 1.993 0 0 0-1 3.72V6.5l3 3v1.78A1.993 1.993 0 0 0 5 15a1.993 1.993 0 0 0 1-3.72V9.5l3-3V4.72A1.993 1.993 0 0 0 8 1zM2 4.2C1.34 4.2.8 3.65.8 3c0-.65.55-1.2 1.2-1.2.65 0 1.2.55 1.2 1.2 0 .65-.55 1.2-1.2 1.2zm3 10c-.66 0-1.2-.55-1.2-1.2 0-.65.55-1.2 1.2-1.2.65 0 1.2.55 1.2 1.2 0 .65-.55 1.2-1.2 1.2zm3-10c-.66 0-1.2-.55-1.2-1.2 0-.65.55-1.2 1.2-1.2.65 0 1.2.55 1.2 1.2 0 .65-.55 1.2-1.2 1.2z"
                      ></path>
                    </svg>
                    <p>{repo.forkCount}</p>
                  </span>
                  <span>
                    <svg
                      aria-hidden="true"
                      className="octicon"
                      height="20"
                      role="img"
                      viewBox="0 0 14 16"
                      width="14"
                      fill="rgb(106, 115, 125)"
                      className="repo-star-svg"
                    >
                      <path
                        fill-rule="evenodd"
                        d="M14 6l-4.9-.64L7 1 4.9 5.36 0 6l3.6 3.26L2.67 14 7 11.67 11.33 14l-.93-4.74L14 6z"
                      ></path>
                    </svg>
                    <p>{repo.stargazers.totalCount}</p>
                  </span>
                </div>
              </div>
            </div>
          );
        })}
      </a>
    </StyledRepo>
  );
};

export const StyledRepo = styled.div`
  .repo-card-div {
    color: rgb(88, 96, 105);
    background-color: rgb(255, 255, 255);
    box-shadow: rgba(0, 0, 0, 0.2) 0px 10px 30px -15px;
    padding: 2rem;
    cursor: pointer;
    transition: all 0.3s;
  }
  .repo-card-div:hover {
    box-shadow: rgba(0, 0, 0, 0.2) 0px 20px 30px -10px;
    transition: all 0.3s;
  }

  .repo-stats {
    display: flex;
    -webkit-box-pack: justify;
    justify-content: space-between;
    font-size: 13px;
    color: rgb(106, 115, 125);
  }

  .repo-left-stat {
    -webkit-box-flex: 1;
    flex-grow: 1;
    display: flex;
  }

  .repo-left-stat span {
    display: flex;
    -webkit-box-align: center;
    align-items: center;
    margin-right: 0.75rem;
  }

  .repo-name-div {
    display: flex;
    align-items: center;
  }

  .repo-svg {
    margin-right: 0.5rem;
    min-width: 16px;
  }

  .repo-name {
    white-space: nowrap;
    text-overflow: ellipsis;
    color: rgb(36, 41, 46);
    margin-bottom: 0.75rem;
    font-size: 25px;
    font-weight: 700;
    letter-spacing: -0.5px;
    overflow: hidden;
    margin: 0px;
  }

  .repo-star-svg {
    margin-right: 0.3rem;
  }

  .repo-description {
    overflow: hidden;
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
  }
`;

export default ContributedRepos;
