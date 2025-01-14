behavior PrettyDate(date)

	on load

		if date == 0 then 
			exit
		end
			
		repeat forever 
			make a Date from (date) called original
			make a Date from (Date.now()) called now
			
			set milisecondCount to (now - original)
			set secondCount to Math.floor(milisecondCount / 1000)
		
			if secondCount < 60 then
				set my innerHTML to "just now"
				wait ((60 * 1000) - milisecondCount) ms 
				continue
			end
			
			set minuteCount to Math.floor(secondCount / 60)

			if minuteCount < 60 then 
				set my innerHTML to minuteCount + "min ago"
				set delay to (60000 - Math.floor(milisecondCount / 60000))
				wait delay ms
				continue
			end

			set hourCount to Math.floor(minuteCount / 60)

			if hourCount < 24 then
				set my innerHTML to hourCount + "h ago"
				exit
			end

			set dayCount to Math.floor(hourCount / 24)

			if dayCount < 4 then
				set my innerHTML to dayCount + "d ago"
				exit
			end

			set my innerHTML to original.toLocaleDateString('en-US', {
				day:'numeric',
				month:'long',
				year:'numeric'
			}) + " &middot; " + original.toLocaleTimeString('en-US', {
				hour12: true,
				hour:'numeric',
				minute: '2-digit'
			})

		end
	end
end