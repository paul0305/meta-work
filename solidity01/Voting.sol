// SPDX-License-Identifier: MIT
pragma solidity ^0.8.2;

contract Voting{
    address public owner;

    mapping (string => uint) public  voteNum;
    //候选人
    string[] public candidates;
    //用户给候选人投票
    function vote(string calldata name)public{
        //string memory pname = name;
       bool exist = false;
       //比较数组是是否存在候选者 
       for(uint i=0;i<candidates.length;i++){
        if(keccak256(abi.encodePacked(candidates[i])) == keccak256(abi.encodePacked(name))){
            exist = true;
            break ;
        }
       }
      if(!exist){
        candidates.push(name);
      }
       voteNum[name] +=1;
    
    }

    //查看票数
    function getVotes(string calldata name)public view returns (uint){
        return voteNum[name];
    }

    //重置所有候选人的得票数
    function resetVotes()public {
        for(uint i = 0;i<candidates.length;i++){
            voteNum[candidates[i]]=0;
        }
        
    }

}
