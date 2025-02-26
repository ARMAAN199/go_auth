// import data from ./Untitled3.js
// search recurively for in the data object to any depth and find all arrays named 'edges'
// each array named 'edges' will have objects with a key 'next_node_id' and 'rule'
// print duplicates from the 'edges' where both 'next_node_id' and 'rule' are the same

import data from "./Untitled3.js";

const removeDuplicates = (data) => {
  const result = [];
  const findDuplicates = (data) => {
    if (Array.isArray(data)) {
      data.forEach((item) => {
        if (item.edges) {
          const duplicates = {};
          item.edges.forEach((edge) => {
            const key = `${edge.next_node_id}-${edge.rule}`;
            if (duplicates[key]) {
              result.push(edge);
            } else {
              duplicates[key] = true;
            }
          });
        } else {
          findDuplicates(item);
        }
      });
    } else if (typeof data === "object") {
      Object.keys(data).forEach((key) => {
        findDuplicates(data[key]);
      });
    }
  };
  findDuplicates(data);
  return result;
};

console.log(removeDuplicates(data));
