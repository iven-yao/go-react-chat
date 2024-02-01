var bcrypt = require('bcryptjs')
var salt = bcrypt.genSaltSync(10);

export const activeTab = "p-2 bg-black text-white rounded-t-xl border-black";
export const inactiveTab = "p-2 rounded-t-xl bg-gray-300  border-b-0 text-gray-600";

export const compare = (hashed: string, pwd: string) => {
    return bcrypt.compareSync(pwd, hashed);
}