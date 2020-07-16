import GithubProfile from "../GithubProfile";

export default function Header() {
  return (
    <section class="landing" id="top">
      <div class="row vertical-center">
          <div class="col-lg-8 text-left pb-5">
              <h1 class="display-1 section-heading">Sebastian Crossa</h1>
                <h1 class="landing-bg-text bg-text">Sebastian Crossa</h1>
              <p class="lead pl-2">Full Stack Developer</p>
          </div>
          <div class="col-lg-4 text-center pb-5">
            <img className="img-fluid headerImg mb-3" style={{ margin: "0 auto" }} src="/profile.jpg" alt="User profile pic"/>
            <GithubProfile />
          </div>
      </div>
    </section>
  );
}
