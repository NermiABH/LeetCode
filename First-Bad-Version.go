/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    left, right := 1, n-1
    lastBad := n
    for left <= right {
        mid := left + (right - left) / 2
        if isBadVersion(mid){
            lastBad=mid
            right = mid-1
        }else {
            left = mid+1
        }
    }
    return lastBad
}