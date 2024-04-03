export class ILogin {
    username: string;
    password: string;
}

export class User  {
  IdUser: number;
  Username: string;
  Password: string;

  // store hashed in database..
  // AccountCode string;
  // LastPasswordID: number;
}
