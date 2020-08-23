export const PodMember = ({ podMemberInfo }) => {
  return (
    <StyledRepo>
      <div className="profile">
        <div className="image">
          <div className="circle-1"></div>
          <div className="circle-2"></div>
          <img src="https://picsum.photos/128" width="70" height="70" />
        </div>

        <div className="name">[Podmate Name]</div>

        <div className="actions">
          <a className="btn">GitHub</a>
        </div>
      </div>
    </StyledRepo>
  );
};

export const StyledRepo = styled.div`
  .profile {
    float: left;
    width: 200px;
    height: 320px;
    text-align: center;
  }
  .profile .image {
    position: relative;
    width: 70px;
    height: 70px;
    margin: 38px auto 0 auto;
  }
  .profile .image .circle-1,
  .profile .image .circle-2 {
    position: absolute;
    box-sizing: border-box;
    width: 76px;
    height: 76px;
    top: -3px;
    left: -3px;
    border-width: 1px;
    border-style: solid;
    border-color: #222222 #222222 #222222 transparent;
    border-radius: 50%;
    -webkit-transition: all 1.5s ease-in-out;
    transition: all 1.5s ease-in-out;
  }
  .profile .image .circle-2 {
    width: 82px;
    height: 82px;
    top: -6px;
    left: -6px;
    border-color: #222222 transparent #222222 #222222;
  }
  .profile .image img {
    display: block;
    border-radius: 50%;
    background: #f5e8df;
  }
  .profile .image:hover {
    cursor: pointer;
  }
  .profile .image:hover .circle-1,
  .profile .image:hover .circle-2 {
    -webkit-transform: rotate(360deg);
    transform: rotate(360deg);
  }
  .profile .image:hover .circle-2 {
    -webkit-transform: rotate(-360deg);
    transform: rotate(-360deg);
  }
  .profile .name {
    font-size: 15px;
    font-weight: 600;
    margin-top: 20px;
  }
  .profile .job {
    font-size: 11px;
    line-height: 15px;
  }
  .profile .actions {
    margin-top: 33px;
  }
  .profile .actions .btn {
    display: block;
    width: 120px;
    height: 30px;
    margin: 0 auto 10px auto;
    background: none;
    border: 1px solid #222222;
    border-radius: 15px;
    font-size: 12px;
    font-weight: 600;
    box-sizing: border-box;
    -webkit-transition: all 0.3s ease-in-out;
    transition: all 0.3s ease-in-out;
    color: #222222;
  }
  .profile .actions .btn:hover {
    background: #f8b92a;
    border: 1px solid #f8b92a;
    color: #fff;
  }
`;
