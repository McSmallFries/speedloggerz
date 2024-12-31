
export class User  {
  IdUser: number;
  Username: string;
  Password: string;

  // store hashed in database (UserPasswords table: [IdUser, IdPassword, Password string, ChangeDate])..
  // AccountCode string;
  // LastPasswordID: number;
}
