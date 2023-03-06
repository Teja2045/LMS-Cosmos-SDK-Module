class Solution { 
public: 
    int findValidSplit(vector<int>& nums) { 
        int n=nums.size(); 
        map<int,int>map1; 
        for(int i=n-1;i>=0;i--){ 
            int current=nums[i]; 
            for(int j=2;j<=sqrt(nums[i]);j++){ 
                if(current%j==0){ 
                    if(map1[j]==0){ 
                        map1[j]=i+1; 
                    } 
                    while(current%j==0) 
                        current/=j; 
                } 
            } 
            if(current>1){ 
                if(map1[current]==0){ 
                        map1[current]=i+1; 
                    } 
            } 
        } 
        int index=1; 
        for(int i=0;i<index;i++){ 
            int current=nums[i]; 
            for(int j=2;j<=sqrt(nums[i]);j++){ 
                if(current%j==0){ 
                    index=max(index,map1[j]); 
                } 
                while(current%j==0) 
                        current/=j; 
            } 
            if(current>1){ 
                index=max(index,map1[current]); 
            } 
        } 
        if(index==n) 
            return -1; 
        return index - 1; 
         
    } 
};