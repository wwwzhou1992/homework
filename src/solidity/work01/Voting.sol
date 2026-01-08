pragma solidity ^0.8.0;

/**
 * @title Voting 投票合约
 * @dev 实现了一个简单的投票系统，允许用户为候选人投票、查询票数和重置投票
 */
contract Voting {
    // 存储候选人得票数：使用string类型作为键（候选人名称），uint256类型作为值（得票数）
    mapping(string => uint256) private votes;
    
    // 存储所有候选人的数组，用于重置操作和获取候选人列表
    string[] private candidates;
    
    // 投票事件：当用户投票时触发
    // @param voter 投票人地址
    // @param candidate 候选人名称
    // @param newVoteCount 候选人的新得票数
    event Voted(address indexed voter, string candidate, uint256 newVoteCount);
    
    // 重置事件：当重置所有投票时触发
    // @param resetter 执行重置操作的地址
    event VotesReset(address indexed resetter);
    
    /**
     * @dev 投票给某个候选人
     * @param candidate 候选人名称
     */
    function vote(string memory candidate) public {
        // 检查候选人名称是否为空
        require(bytes(candidate).length > 0, "Candidate name cannot be empty");
        
        // 如果是第一次给该候选人投票，将其添加到候选人列表中
        if (votes[candidate] == 0 && !_isCandidateInList(candidate)) {
            candidates.push(candidate);
        }
        
        // 给候选人的得票数加1
        votes[candidate]++;
        
        // 触发投票事件
        emit Voted(msg.sender, candidate, votes[candidate]);
    }
    
    /**
     * @dev 获取候选人的得票数
     * @param candidate 候选人名称
     * @return 该候选人的得票数
     */
    function getVotes(string memory candidate) public view returns (uint256) {
        // 返回指定候选人的得票数
        return votes[candidate];
    }
    
    /**
     * @dev 重置所有候选人的得票数
     */
    function resetVotes() public {
        // 遍历所有候选人，删除他们的得票数
        for (uint256 i = 0; i < candidates.length; i++) {
            delete votes[candidates[i]];
        }
        
        // 清空候选人列表
        delete candidates;
        
        // 触发重置事件
        emit VotesReset(msg.sender);
    }
    
    /**
     * @dev 内部函数：检查候选人是否已在列表中
     * @param candidate 候选人名称
     * @return 是否在列表中
     */
    function _isCandidateInList(string memory candidate) private view returns (bool) {
        // 遍历候选人列表，使用keccak256哈希比较字符串是否相等
        for (uint256 i = 0; i < candidates.length; i++) {
            if (keccak256(abi.encodePacked(candidates[i])) == keccak256(abi.encodePacked(candidate))) {
                return true;
            }
        }
        return false;
    }
    
    /**
     * @dev 获取所有候选人
     * @return 候选人数组
     */
    function getCandidates() public view returns (string[] memory) {
        // 返回所有候选人的数组
        return candidates;
    }
}