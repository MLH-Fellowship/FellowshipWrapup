import GithubProfile from "../GithubProfile";

export default function Header({ accountInfo }) {
  const { Name, AvatarUrl, Company } = accountInfo.User;

  return (
    <section className="landing" id="top">
      <div className="row vertical-center">
        <div className="col-lg-8 text-left pb-5">
          <h1 className="display-1 section-heading">{Name}</h1>
          <h1 className="landing-bg-text bg-text">{Name}</h1>
          <p className="lead pl-2">{Company}</p>
        </div>
        <div className="col-lg-4 text-center pb-5">
          <img
            className="img-fluid headerImg mb-3"
            style={{ margin: "0 auto" }}
            src={`${AvatarUrl}`}
            alt="User profile pic"
          />
          <GithubProfile accountInfo={accountInfo.User} />
        </div>
      </div>
    </section>
  );
}
