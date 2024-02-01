var bcrypt = require('bcryptjs')
var salt = bcrypt.genSaltSync(10);

export const compare = (hashed: string, pwd: string) => {
    return bcrypt.compareSync(pwd, hashed);
}